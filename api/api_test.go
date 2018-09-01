package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albshin/tekkenmatchdb-backend/store"
	"github.com/go-chi/chi"
)

type MockDB struct {
	store.Store
}

func testJSON(srv *chi.Mux, method, endpoint, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, endpoint, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	return rec
}

func newServer(t *testing.T) *chi.Mux {
	db, err := store.Open("user=postgres password=password dbname=matches sslmode=disable")
	if err != nil {
		t.Fatalf("Could not initalize the server.\n")
	}
	h := Handler{Store: db}
	return NewRouter(&h)
}
