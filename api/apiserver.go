package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/liftM/pooter/effects/pooterdb"
)

type Server struct {
	Router *chi.Mux
	DB     *pooterdb.Postgres
}

func NewServer(conn string) *Server {
	var err error
	s := Server{}

	// Connect to dependencies.
	s.DB, err = pooterdb.New(context.Background(), conn)
	if err != nil {
		panic(err)
	}

	// Set up router.
	s.Router = chi.NewRouter()

	// Set up middleware.
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// Register routes.
	s.Router.Handle("/users.create", http.HandlerFunc(s.CreateUser))
	s.Router.Handle("/users.follow", http.HandlerFunc(s.FollowUser))
	s.Router.Handle("/users.posts", http.HandlerFunc(s.ListUserPosts))
	s.Router.Handle("/poots.post", http.HandlerFunc(s.CreatePost))

	return &s
}
