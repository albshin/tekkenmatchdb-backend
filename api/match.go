package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/albshin/tekkenmatchdb-backend/model"
	"github.com/go-chi/chi"
)

func (h *Handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	pageParams, err := withPagination(r)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	matchFilters, err := getMatchFilters(r)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Store.GetMatches(matchFilters, pageParams)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}

func (h *Handler) GetMatchesByPlayerID(w http.ResponseWriter, r *http.Request) {
	pageParams, err := withPagination(r)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	matchFilters, err := getMatchFilters(r)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	playerID, err := strconv.Atoi(chi.URLParam(r, "player_id"))
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Store.GetMatchesByPlayerID(playerID, matchFilters, pageParams)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}

func (h *Handler) CreateMatches(w http.ResponseWriter, r *http.Request) {
	var req []*model.Match
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "error", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.Store.CreateMatches(req)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusCreated)
}
