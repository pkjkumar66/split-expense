package handlers

import (
	"net/http"

	"splitexpense/internal/services"
	"splitexpense/internal/utils"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
	balanceService *services.BalanceService
}

func NewExpenseHandler(expenseService *services.ExpenseService, balanceService *services.BalanceService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
		balanceService: balanceService,
	}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *ExpenseHandler) ListExpenses(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
