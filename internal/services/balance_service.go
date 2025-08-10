package services

import (
	"database/sql"
)

type BalanceService struct {
	db *sql.DB
}

func NewBalanceService(db *sql.DB) *BalanceService {
	return &BalanceService{db: db}
}
