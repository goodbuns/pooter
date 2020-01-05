package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req CreatePostRequest
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

	// Create post.
	err = s.DB.CreatePost(ctx, req.UserID, req.Content)
	if err != nil {
		panic(err)
	}

	// Return created post content.
	res, err := json.Marshal(CreatePostResponse{Content: req.Content})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}
