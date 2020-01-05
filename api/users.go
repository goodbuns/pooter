package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/liftM/pooter/effects/pooterdb"
)

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	UserID pooterdb.UserID `json:"user_id"`
}

type FollowUserRequest struct {
	UserID   string `json:"user_id"`
	FollowID string `json:"follow_id"`
}

type FollowUserResponse struct {
	UserID string `json:"user_id"`
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
	uid, err := s.DB.CreateUser(ctx, req.Username, req.Password)
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
	if err := s.DB.FollowUser(ctx, req.UserID, req.FollowID); err != nil {
		panic(err)
	}

	// Return ID of followed user.
	res, err := json.Marshal(FollowUserResponse{UserID: req.FollowID})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}
