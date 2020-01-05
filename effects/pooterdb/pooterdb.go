package pooterdb

import (
	"context"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type UserID string

type PooterDB interface {
	CreateUser(ctx context.Context, username, password string) (UserID, error)
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

func (p *Postgres) CreateUser(ctx context.Context, username, password string) (UserID, error) {
	result := p.db.QueryRowContext(ctx,
		`INSERT INTO users
			(id, username, password)
		VALUES
			(DEFAULT, $1, $2)
		RETURNING id`, username, password)

	var id int
	err := result.Scan(&id)
	if err != nil {
		return "", err
	}
	return UserID(strconv.Itoa(id)), nil
}

func (p *Postgres) FollowUser(ctx context.Context, userID, followID string) error {
	u, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	f, err := strconv.Atoi(followID)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(
		`INSERT INTO followers
			(id, userid, followid)
		VALUES
			(DEFAULT, $1, $2)
		RETURNING id`, u, f)

	if err != nil {
		return err
	}
	return nil
}
