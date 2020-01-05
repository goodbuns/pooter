package api

import (
	"net/http"

	"github.com/liftM/pooter/types"
)

type CreateUserRequest struct {
	Username string
	Password string
}

type FollowUserRequest struct {
	Username string
	Idol     string
}

type ListUserPostsRequest struct {
	UserID types.UserID `json:"user_id"`
}

type ListUserPostsResponse struct {
	Posts []types.Post
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	var req CreateUserRequest
	s.ReadRequest(r, &req)

	// Create user.
	if err := s.db.CreateUser(ctx, req.Username, req.Password); err != nil {
		panic(err)
	}

	// Return empty response.
	_, err := w.Write([]byte{})
	if err != nil {
		panic(err)
	}
}

func (s *Server) FollowUser(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	var req FollowUserRequest
	s.ReadRequest(r, &req)

	// Follow user.
	if err := s.db.FollowUser(ctx, req.Username, req.Idol); err != nil {
		panic(err)
	}

	// Return empty response.
	_, err := w.Write([]byte{})
	if err != nil {
		panic(err)
	}
}

func (s *Server) ListUserPosts(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	var req ListUserPostsRequest
	s.ReadRequest(r, &req)

	// Retrieve all posts from particular user.
	p, err := s.db.ListUserPosts(ctx, req.UserID)
	if err != nil {
		panic(err)
	}

	// Return all posts for particular user.
	resp := ListUserPostsResponse{Posts: p}
	s.WriteResponse(w, &resp)
}
