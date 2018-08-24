package api

import (
	"encoding/json"
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/store"
)

type Handler struct {
	Store *store.PGStore
}

/*
type LoginPost struct {
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginPost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "error", http.StatusInternalServerError)
	}
	defer r.Body.Close()

	expires := time.Now().Add(time.Hour * 24 * 365)
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   "1",
		ExpiresAt: expires.Unix(),
	}).SignedString([]byte("SECRET"))
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenStr,
		Path:    "/",
		Expires: expires,
	})
	// TODO: Send response back with JWT
}
*/

func sendJSON(w http.ResponseWriter, resp interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendError(w http.ResponseWriter, msg string, code int) {
	sendJSON(w, map[string]string{"error": msg}, code)
}
