package services

import (
	"database/sql"
	"splitexpense/internal/models"
)

type ExpenseService struct {
	db *sql.DB
}

func NewExpenseService(db *sql.DB) *ExpenseService {
	return &ExpenseService{db: db}
}

func (s *ExpenseService) CreateExpense(req models.CreateExpenseRequest) (*models.Expense, error) {
	// In a real application, you would have more complex logic here,
	// such as validating the shares, checking user permissions, etc.
	expense := &models.Expense{
		Amount:      req.Amount,
		Currency:    req.Currency,
		Description: req.Description,
	}
	// This is a placeholder. You would need to insert the expense and shares into the database.
	return expense, nil
}
