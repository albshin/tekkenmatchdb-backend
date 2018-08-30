package api

import "net/http"

func (h *Handler) GetRanks(w http.ResponseWriter, r *http.Request) {
	res, err := h.Store.GetRanks()
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}

func (h *Handler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	res, err := h.Store.GetCharacters()
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, &res, http.StatusOK)
}
