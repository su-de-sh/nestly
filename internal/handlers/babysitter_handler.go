package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/su-de-sh/nestly/internal/api/request"
	"github.com/su-de-sh/nestly/internal/repository"
)

type BabysitterHandler struct {
	babysitterRepository *repository.BabySitterRepository
}

func NewBabysitterHandler(babysitterRepository *repository.BabySitterRepository,
) *BabysitterHandler {
	return &BabysitterHandler{
		babysitterRepository: babysitterRepository,
	}
}

func (h *BabysitterHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req request.BabysitterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.babysitterRepository.CreateBabySitter(req)

	if err != nil {
		http.Error(w, "Error creating babysitter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *BabysitterHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	babySitters := h.babysitterRepository.GetBabySitters()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babySitters)

}

func (h *BabysitterHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var req request.BabysitterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request error", http.StatusBadRequest)
		return
	}

	if err := h.babysitterRepository.UpdateById(id, req); err != nil {
		http.Error(w, "Error updating babysitter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *BabysitterHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL parameter
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	if err := h.babysitterRepository.DeleteById(id); err != nil {
		http.Error(w, "Error deleting babysitter", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
