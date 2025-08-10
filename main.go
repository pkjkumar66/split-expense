package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"splitexpense/internal/config"
	"splitexpense/internal/database"
	"splitexpense/internal/handlers"
	"splitexpense/internal/middleware"
	"splitexpense/internal/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize config
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize services
	authService := services.NewAuthService(db, cfg.JWTSecret)
	userService := services.NewUserService(db)
	groupService := services.NewGroupService(db)
	expenseService := services.NewExpenseService(db)
	balanceService := services.NewBalanceService(db)
	settlementService := services.NewSettlementService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	groupHandler := handlers.NewGroupHandler(groupService)
	expenseHandler := handlers.NewExpenseHandler(expenseService, balanceService)
	balanceHandler := handlers.NewBalanceHandler(balanceService)
	settlementHandler := handlers.NewSettlementHandler(settlementService, balanceService)

	// Setup router
	r := mux.NewRouter()

	// Add middleware
	r.Use(middleware.CORS)
	r.Use(middleware.Logger)

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", authHandler.SignUp).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
	auth.HandleFunc("/refresh", authHandler.RefreshToken).Methods("POST")
	auth.HandleFunc("/forgot-password", authHandler.ForgotPassword).Methods("POST")
	auth.HandleFunc("/reset-password", authHandler.ResetPassword).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("/").Subrouter()
	protected.Use(func(next http.Handler) http.Handler {
		return middleware.AuthRequired(authService, next)
	})

	// User routes
	protected.HandleFunc("/users/me", userHandler.GetMe).Methods("GET")
	protected.HandleFunc("/users/me", userHandler.UpdateMe).Methods("PUT")

	// Group routes
	groups := protected.PathPrefix("/groups").Subrouter()
	groups.HandleFunc("", groupHandler.CreateGroup).Methods("POST")
	groups.HandleFunc("", groupHandler.ListGroups).Methods("GET")
	groups.HandleFunc("/{id}", groupHandler.GetGroup).Methods("GET")
	groups.HandleFunc("/{id}", groupHandler.UpdateGroup).Methods("PUT")
	groups.HandleFunc("/{id}", groupHandler.DeleteGroup).Methods("DELETE")
	groups.HandleFunc("/{id}/invite", groupHandler.InviteToGroup).Methods("POST")
	groups.HandleFunc("/{id}/join", groupHandler.JoinGroup).Methods("POST")
	groups.HandleFunc("/{id}/leave", groupHandler.LeaveGroup).Methods("DELETE")

	// Expense routes
	expenses := protected.PathPrefix("/expenses").Subrouter()
	expenses.HandleFunc("", expenseHandler.CreateExpense).Methods("POST")
	expenses.HandleFunc("", expenseHandler.ListExpenses).Methods("GET")
	expenses.HandleFunc("/{id}", expenseHandler.GetExpense).Methods("GET")
	expenses.HandleFunc("/{id}", expenseHandler.UpdateExpense).Methods("PUT")
	expenses.HandleFunc("/{id}", expenseHandler.DeleteExpense).Methods("DELETE")

	// Balance routes
	protected.HandleFunc("/balances/{user_id}", balanceHandler.GetUserBalances).Methods("GET")
	protected.HandleFunc("/groups/{id}/balances", balanceHandler.GetGroupBalances).Methods("GET")
	protected.HandleFunc("/groups/{id}/simplify", balanceHandler.SimplifyDebts).Methods("POST")

	// Settlement routes
	settlements := protected.PathPrefix("/settlements").Subrouter()
	settlements.HandleFunc("", settlementHandler.CreateSettlement).Methods("POST")
	settlements.HandleFunc("", settlementHandler.ListSettlements).Methods("GET")
	settlements.HandleFunc("/{id}", settlementHandler.GetSettlement).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
