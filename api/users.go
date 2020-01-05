package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/liftM/pooter/types"
)

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	UserID types.UserID `json:"user_id"`
}

type FollowUserRequest struct {
	UserID         types.UserID `json:"user_id"`
	FollowedUserID types.UserID `json:"follow_id"`
}

type FollowUserResponse struct {
	UserID types.UserID `json:"user_id"`
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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req CreateUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	// Create user.
	uid, err := s.db.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		panic(err)
	}

	// Return ID of created user.
	res, err := json.Marshal(CreateUserResponse{UserID: uid})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}

func (s *Server) FollowUser(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req FollowUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	// Follow user.
	if err := s.db.FollowUser(ctx, req.UserID, req.FollowedUserID); err != nil {
		panic(err)
	}

	// Return ID of followed user.
	res, err := json.Marshal(FollowUserResponse{UserID: req.FollowedUserID})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}

func (s *Server) ListUserPosts(w http.ResponseWriter, r *http.Request) {
	// Read request.
	ctx := r.Context()

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req ListUserPostsRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	// Retrieve all posts from particular user.
	p, err := s.db.ListUserPosts(ctx, req.UserID)
	if err != nil {
		panic(err)
	}

	// Return all posts for particular user.
	res, err := json.Marshal(ListUserPostsResponse{Posts: p})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}
