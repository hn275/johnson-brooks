package server

import (
	"jb/mod/ping"

	"github.com/go-chi/chi/v5"
)

func RouteServer(r *chi.Mux) {
	r.Handle("/ping", ping.Controller())
}
