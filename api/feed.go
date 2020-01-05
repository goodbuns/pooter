package api

import (
	"errors"
	"net/http"

	"github.com/liftM/pooter/types"
)

type ViewFeedRequest struct {
	UserID   types.UserID `json:"user_id"`
	Password string
	Page     int
	Limit    int
}

type ViewFeedResponse struct {
	Posts []types.Post
}

func (s *Server) ViewFeed(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	var req ViewFeedRequest
	s.ReadRequest(r, &req)

	// Verify auth.
	ok, err := s.VerifyAuth(ctx, req.UserID, req.Password)
	if err != nil {
		panic(err)
	} else if !ok {
		panic(errors.New("incorrect username or password given"))
	}

	// View feed.
	posts, err := s.db.ViewFeed(ctx, req.UserID, req.Page, req.Limit)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	resp := ViewFeedResponse{Posts: posts}
	s.WriteResponse(w, &resp)
}
