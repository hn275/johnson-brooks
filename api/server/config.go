package server

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var appEnv string

func init() {

	appEnv = os.Getenv("APP_ENV")
}

func ConfigServer(r *chi.Mux) {
	if appEnv == "development" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)
}
