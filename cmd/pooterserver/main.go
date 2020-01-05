package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/liftM/pooter/api"
)

func main() {
	// Get configuration.
	conn := flag.String("db", "", "database connection string")
	flag.Parse()

	s := api.NewServer(*conn)

	// r.Handle("/users.posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// r.Handle("/poots.post", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// r.Handle("/poots.feed", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", s.Router)
}
