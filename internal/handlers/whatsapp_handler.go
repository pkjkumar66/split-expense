package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"splitexpense/internal/config"
	"splitexpense/internal/models"
	"splitexpense/internal/services"
	"splitexpense/internal/utils"
)

type WhatsAppHandler struct {
	whatsAppService *services.WhatsAppService
	cfg             *config.Config
}

func NewWhatsAppHandler(whatsAppService *services.WhatsAppService, cfg *config.Config) *WhatsAppHandler {
	return &WhatsAppHandler{
		whatsAppService: whatsAppService,
		cfg:             cfg,
	}
}

func (h *WhatsAppHandler) Webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		verifyToken := r.URL.Query().Get("hub.verify_token")
		if verifyToken == h.cfg.WhatsAppVerifyToken {
			challenge := r.URL.Query().Get("hub.challenge")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			utils.RespondWithError(w, http.StatusForbidden, "Invalid verify token")
		}
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	var msg models.WhatsAppMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse webhook payload")
		return
	}

	if err := h.whatsAppService.ProcessMessage(msg); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to process message")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Webhook received")
}
