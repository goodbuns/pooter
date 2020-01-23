package main

import (
	"context"
	"flag"
	"log"

	"github.com/liftM/pooter/api"
	"github.com/liftM/pooter/effects/pooterdb"
)

func main() {
	// Get configuration.
	conn := flag.String("db", "", "database connection string")
	flag.Parse()

	db, err := pooterdb.New(context.Background(), *conn)
	if err != nil {
		panic(err)
	}
	s := api.New(db)

	log.Println("Listening on port :8000")
	s.Start()
}
