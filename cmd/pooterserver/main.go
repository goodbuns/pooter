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

	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", s.Router)
}
