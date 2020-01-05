package api

import (
	"errors"
	"net/http"

	"github.com/liftM/pooter/types"
)

type CreatePostRequest struct {
	Content  string
	UserID   types.UserID `json:"user_id"`
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
	ok, err := s.VerifyAuth(ctx, req.UserID, req.Password)
	if err != nil {
		panic(err)
	} else if !ok {
		panic(errors.New("incorrect username or password given"))
	}

	// Create post.
	err = s.db.CreatePost(ctx, req.UserID, req.Content)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	resp := CreatePostResponse{Content: req.Content}
	s.WriteResponse(w, &resp)
}
