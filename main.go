package main

import (
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/api"
	"github.com/albshin/tekkenmatchdb-backend/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

		r.Get("/players", h.GetPlayers)
		r.Get("/players/{player_id}", h.GetPlayer)
		r.Get("/players/{player_id}/matches", h.GetMatchesByPlayerID)
		r.With(jsonRequired).Post("/player", h.CreatePlayer)
	})

	http.ListenAndServe(":3000", r)
}
