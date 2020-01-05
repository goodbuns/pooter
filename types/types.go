package types

import "time"

type Post struct {
	Content   string
	UserID    UserID    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserID string
