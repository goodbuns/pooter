package pooterdb

import (
	"context"
	"strconv"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/liftM/pooter/types"
)

type PooterDB interface {
	CreateUser(ctx context.Context, username, password string) error
	FollowUser(ctx context.Context, user, idol string) error
	RetrievePassword(ctx context.Context, userID types.UserID) (string, error)
	CreatePost(ctx context.Context, userID types.UserID, content string) error
	ListUserPosts(ctx context.Context, userID types.UserID) ([]types.Post, error)
	ViewFeed(ctx context.Context, userID types.UserID, page, limit int) ([]types.Post, error)
}

var _ PooterDB = &Postgres{}

type Postgres struct {
	db *sqlx.DB
}

func New(ctx context.Context, conn string) (*Postgres, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", conn)
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) CreateUser(ctx context.Context, username, password string) error {
	_, err := p.db.Exec(
		`INSERT INTO users
			(id, username, password)
		VALUES
			(DEFAULT, $1, $2)`, username, password)

	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UserID(ctx context.Context, username string) (int, error) {
	var uid int

	result := p.db.QueryRowContext(ctx,
		`SELECT id FROM users
		WHERE username = $1`, username)
	err := result.Scan(&uid)
	if err != nil {
		return uid, err
	}

	return uid, nil
}

func (p *Postgres) FollowUser(ctx context.Context, username, idol string) error {
	u, err := p.UserID(ctx, username)
	if err != nil {
		return err
	}

	i, err := p.UserID(ctx, idol)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(
		`INSERT INTO followers
			(id, user_id, idol)
		VALUES
			(DEFAULT, $1, $2)
		RETURNING id`, u, i)

	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) RetrievePassword(ctx context.Context, userID types.UserID) (string, error) {
	var password string
	u, err := strconv.Atoi(string(userID))
	if err != nil {
		return password, err
	}

	result := p.db.QueryRowContext(ctx,
		`SELECT password FROM users
		WHERE id = $1`, u)

	err = result.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (p *Postgres) CreatePost(ctx context.Context, userID types.UserID, content string) error {
	u, err := strconv.Atoi(string(userID))
	if err != nil {
		return err
	}

	_, err = p.db.Exec(
		`INSERT INTO posts
			(id, content, author, created_at)
		VALUES
			(DEFAULT, $1, $2, NOW())
		RETURNING id`, content, u)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) ListUserPosts(ctx context.Context, userID types.UserID) ([]types.Post, error) {
	var posts []types.Post
	u, err := strconv.Atoi(string(userID))
	if err != nil {
		return posts, err
	}
	rows, err := p.db.QueryContext(ctx,
		`SELECT content, author, created_at FROM posts
		WHERE user_id = $1`, u)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var user int64
		var content string
		var createdAt time.Time
		if err := rows.Scan(&content, &user, &createdAt); err != nil {
			return posts, err
		}
		posts = append(posts, types.Post{Content: content, UserID: types.UserID(string(int(user))), CreatedAt: createdAt})
	}
	return posts, nil
}

func (p *Postgres) ViewFeed(ctx context.Context, userID types.UserID, page, limit int) ([]types.Post, error) {
	var posts []types.Post
	u, err := strconv.Atoi(string(userID))
	if err != nil {
		return posts, err
	}

	// Find all users the user is following.
	var followingUsers []string
	rows, err := p.db.QueryContext(ctx,
		`SELECT idol
		FROM users INNER JOIN followers
		ON users.id = followers.user_id AND
		users.id = $1`, u)

	defer rows.Close()
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return posts, err
		}
		followingUsers = append(followingUsers, user)
	}

	// Find most recent posts.
	offset := page * limit
	if offset < 1 {
		offset = 0
	}

	param := "{" + strings.Join(followingUsers, ",") + "}"
	postRows, err := p.db.QueryContext(ctx,
		`SELECT content, author, created_at FROM posts
		WHERE author = ANY($1::int[])
		ORDER BY id DESC
		LIMIT $2
		OFFSET $3`, param, limit, offset)
	if err != nil {
		return posts, err
	}
	defer postRows.Close()

	for postRows.Next() {
		var user int64
		var content string
		var createdAt time.Time
		if err := postRows.Scan(&content, &user, &createdAt); err != nil {
			return posts, err
		}
		posts = append(posts, types.Post{Content: content, UserID: types.UserID(string(int(user))), CreatedAt: createdAt})
	}
	return posts, nil
}
