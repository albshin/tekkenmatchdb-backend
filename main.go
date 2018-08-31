package main

import (
	"net/http"

	"github.com/albshin/tekkenmatchdb-backend/api"
	"github.com/albshin/tekkenmatchdb-backend/store"
	_ "github.com/lib/pq"
)

func main() {
	db, err := store.Open("user=postgres password=password dbname=matches sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	h := api.Handler{Store: db}
	r := api.NewRouter(&h)

	http.ListenAndServe(":4000", r)
}
