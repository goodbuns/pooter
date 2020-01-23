package api

import (
	"errors"
	"net/http"
)

type CreatePostRequest struct {
	Content  string
	Username string
	Password string
}

type CreatePostResponse struct {
	Content string
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	var req CreatePostRequest
	s.ReadRequest(r, &req)

	// Verify auth.
	ok, err := s.db.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		panic(err)
	} else if !ok {
		panic(errors.New("incorrect username or password given"))
	}

	// Create post.
	err = s.db.CreatePost(ctx, req.Username, req.Content)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	resp := CreatePostResponse{Content: req.Content}
	s.WriteResponse(w, &resp)
}
