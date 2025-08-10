package services

import (
	"database/sql"
)

type SettlementService struct {
	db *sql.DB
}

func NewSettlementService(db *sql.DB) *SettlementService {
	return &SettlementService{db: db}
}
