package services

import (
	"database/sql"
)

type ExpenseService struct {
	db *sql.DB
}

func NewExpenseService(db *sql.DB) *ExpenseService {
	return &ExpenseService{db: db}
}
