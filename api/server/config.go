package server

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var appEnv string
var clientUri string

func init() {

	clientUri = os.Getenv("CLIENT_URI")
	if clientUri == "" {
		panic("[ERROR] missing env: CLIENT_URI")
	}
	appEnv = os.Getenv("APP_ENV")
}

func ConfigServer(r *chi.Mux) {
	if appEnv == "development" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{clientUri},
	}))

}
