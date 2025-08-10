package handlers

import (
	"net/http"

	"splitexpense/internal/services"
)

type GroupHandler struct {
	groupService *services.GroupService
}

func NewGroupHandler(groupService *services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) ListGroups(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) InviteToGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) JoinGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (h *GroupHandler) LeaveGroup(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
