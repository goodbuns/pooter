package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/liftM/pooter/types"
)

type ViewFeedRequest struct {
	Username   string
	Password   string
	PageSize   int   `json:"page_size"`
	BeforeTime int64 `json:"before_time"`
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

	// Convert Unix Timestamp to time.Time.
	t := time.Unix(req.BeforeTime, 0)

	// View feed.
	posts, err := s.db.ViewFeed(ctx, req.Username, t, req.PageSize)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	resp := ViewFeedResponse{Posts: posts}
	s.WriteResponse(w, &resp)
}
