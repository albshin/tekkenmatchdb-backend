package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/albshin/tekkenmatchdb-backend/model"
)

func (h *Handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	pageParams, err := withPagination(r)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Store.GetPlayers(pageParams)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}

func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	req := &model.Player{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "error", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.Store.CreatePlayer(req)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusCreated)
}

func (h *Handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := chi.URLParam(r, "player_id")
	pID, err := strconv.Atoi(playerID)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.Store.GetPlayer(pID)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}
