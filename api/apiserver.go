package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/liftM/pooter/effects/pooterdb"
)

type Server struct {
	router *chi.Mux
	db     *pooterdb.Postgres
}

func NewServer(db *pooterdb.Postgres) *Server {
	// Set up router.
	r := chi.NewRouter()
	s := Server{router: r, db: db}

	// Set up middleware.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register routes.
	r.Handle("/users.create", http.HandlerFunc(s.CreateUser))
	r.Handle("/users.follow", http.HandlerFunc(s.FollowUser))
	r.Handle("/users.posts", http.HandlerFunc(s.ListUserPosts))
	r.Handle("/poots.post", http.HandlerFunc(s.CreatePost))
	r.Handle("/poots.feed", http.HandlerFunc(s.ViewFeed))

	return &s
}

func (s *Server) Start() {
	http.ListenAndServe(":8000", s.router)
}
