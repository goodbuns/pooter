package api

import (
	"errors"
	"net/http"

	"github.com/liftM/pooter/types"
)

type ViewFeedRequest struct {
	Username string
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
	ok, err := s.db.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		panic(err)
	} else if !ok {
		panic(errors.New("incorrect username or password given"))
	}

	// View feed.
	posts, err := s.db.ViewFeed(ctx, req.Username, req.Page, req.Limit)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	resp := ViewFeedResponse{Posts: posts}
	s.WriteResponse(w, &resp)
}
