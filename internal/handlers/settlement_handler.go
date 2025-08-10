package handlers

import (
	"net/http"

	"splitexpense/internal/services"
)

type SettlementHandler struct {
	settlementService *services.SettlementService
	balanceService    *services.BalanceService
}

func NewSettlementHandler(settlementService *services.SettlementService, balanceService *services.BalanceService) *SettlementHandler {
	return &SettlementHandler{
		settlementService: settlementService,
		balanceService:    balanceService,
	}
}

func (h *SettlementHandler) CreateSettlement(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *SettlementHandler) ListSettlements(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *SettlementHandler) GetSettlement(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
