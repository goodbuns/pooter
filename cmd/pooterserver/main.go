package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {})

	r.Post("/users/{userID}/followers", func(w http.ResponseWriter, r *http.Request) {})

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {})

	r.Get("/users/{userID}/feed", func(w http.ResponseWriter, r *http.Request) {})

	r.Get("/users/{userID}/posts", func(w http.ResponseWriter, r *http.Request) {})

	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", r)
}
