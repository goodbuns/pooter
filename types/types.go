package types

import "time"

type Post struct {
	Content   string
	UserID    UserID    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Auth struct {
	Username string
	Password string
}

type UserID string