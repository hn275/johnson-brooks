package server

import (
	"jb/mod/auth"
	"jb/mod/product"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RouteServer(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Use(cors.Handler(cors.Options{
				AllowedOrigins: []string{clientUri},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			}))

			r.Handle("/login", http.HandlerFunc(auth.Login))
		})
	})

	r.Handle("/", http.HandlerFunc(product.GetProducts))
}
