package server

import (
	"jb/mod/auth"
	"jb/mod/ping"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteServer(r *chi.Mux) {
	r.Handle("/ping", ping.Controller())
	r.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Handle("/register", http.HandlerFunc(auth.Register))
		})
	})
}
