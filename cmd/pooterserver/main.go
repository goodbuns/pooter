package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/liftM/pooter/effects/pooterdb"
)

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	UserID pooterdb.UserID `json:"user_id"`
}

func main() {
	// Get configuration.
	conn := flag.String("db", "", "database connection string")
	flag.Parse()

	// Connect to dependencies.
	db, err := pooterdb.New(context.Background(), *conn)
	if err != nil {
		panic(err)
	}

	// Set up router.
	r := chi.NewRouter()

	// Set up middleware.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up routes.
	r.Handle("/users.create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read request.
		ctx := r.Context()

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var req CreateUserRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			panic(err)
		}

		// Create user.
		uid, err := db.CreateUser(ctx, req.Username, req.Password)
		if err != nil {
			panic(err)
		}

		// Return ID of created user.
		res, err := json.Marshal(CreateUserResponse{UserID: uid})
		if err != nil {
			panic(err)
		}

		_, err = w.Write(res)
		if err != nil {
			panic(err)
		}
	}))

	r.Handle("/users.follow", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	r.Handle("/users.posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	r.Handle("/poots.post", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	r.Handle("/poots.feed", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", r)
}
