package handlers

import (
	"net/http"

	"splitexpense/internal/services"
)

type BalanceHandler struct {
	balanceService *services.BalanceService
}

func NewBalanceHandler(balanceService *services.BalanceService) *BalanceHandler {
	return &BalanceHandler{balanceService: balanceService}
}

func (h *BalanceHandler) GetUserBalances(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *BalanceHandler) GetGroupBalances(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *BalanceHandler) SimplifyDebts(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
