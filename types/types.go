package types

import "time"

type Post struct {
	Content   string
	Username  string
	CreatedAt time.Time `json:"created_at"`
}

type Auth struct {
	Username string
	Password string
}
