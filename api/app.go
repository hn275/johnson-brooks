package main

import (
	"jb/server"
	"log"
	"net/http"
	"os"

	_ "jb/database"

	"github.com/go-chi/chi/v5"
)

var port string

func init() {
	p := os.Getenv("PORT")
	if p == "" {
		port = ":8080"
	} else {
		port = ":" + p
	}
}

func main() {
	r := chi.NewMux()
	server.ConfigServer(r)
	server.RouteServer(r)
	log.Println("PORT" + port)
	http.ListenAndServe(port, r)
}
