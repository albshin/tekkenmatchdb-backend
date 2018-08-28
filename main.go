package main

import (
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/api"
	"github.com/albshin/tekkenmatchdb-backend/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	db, err := store.Open("user=postgres password=password dbname=matches sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	h := api.Handler{Store: db}
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
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
	jsonRequired := middleware.AllowContentType("application/json")

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	r.Route("/api", func(r chi.Router) {
		//r.With(jsonRequired).Post("/login", h.LoginUser)

		r.Get("/matches", h.GetMatches)
		r.With(jsonRequired).Post("/matches", h.CreateMatches)
		r.With(jsonRequired).Post("/matches/report", h.CreateMatchReport)

		r.Get("/players", h.GetPlayers)
		r.Get("/players/{player_id}", h.GetPlayer)
		r.Get("/players/{player_id}/matches", h.GetMatchesByPlayerID)
		r.With(jsonRequired).Post("/player", h.CreatePlayer)
	})

	http.ListenAndServe(":4000", r)
}
