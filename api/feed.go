package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req ViewFeedRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	// Verify auth.
	ok, err := s.VerifyAuth(ctx, req.UserID, req.Password)
	if err != nil {
		panic(err)
	} else if !ok {
		panic(errors.New("incorrect username or password given"))
	}

	// View feed.
	posts, err := s.DB.ViewFeed(ctx, req.UserID, req.Page, req.Limit)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	res, err := json.Marshal(ViewFeedResponse{Posts: posts})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}
