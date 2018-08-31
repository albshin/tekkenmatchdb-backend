package api

import (
	"encoding/json"
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Handler struct {
	Store store.Store
}

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//jsonRequired := middleware.AllowContentType("application/json")

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	r.Route("/api", func(r chi.Router) {
		//r.With(jsonRequired).Post("/login", h.LoginUser)

		r.Get("/matches", h.GetMatches)
		//r.With(jsonRequired).Post("/matches", h.CreateMatches)
		//r.With(jsonRequired).Post("/matches/report", h.CreateMatchReport)

		r.Get("/players", h.GetPlayers)
		r.Get("/players/names", h.GetPlayerNames)
		r.Get("/players/{player_id}", h.GetPlayer)
		r.Get("/players/{player_id}/matches", h.GetMatchesByPlayerID)
		//r.With(jsonRequired).Post("/players", h.CreatePlayer)

		r.Get("/characters", h.GetCharacters)
		r.Get("/ranks", h.GetRanks)
	})
	return r
}

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
