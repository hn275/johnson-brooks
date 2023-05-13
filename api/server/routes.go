package server

import (
	"jb/mod/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteServer(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Handle("/login", http.HandlerFunc(auth.Login))
		})
	})
}
