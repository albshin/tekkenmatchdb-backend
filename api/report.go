package api

import (
	"encoding/json"
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/model"
)

func (h *Handler) CreateMatchReport(w http.ResponseWriter, r *http.Request) {
	var req model.MatchReport
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.Store.CreateMatchReport(&req)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJSON(w, res, http.StatusCreated)
}
