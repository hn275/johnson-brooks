package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteServer(r *chi.Mux) {
	r.Handle("GET", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}
