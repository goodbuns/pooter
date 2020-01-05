package api

import (
	"context"

	"github.com/liftM/pooter/types"
)

// VerifyAuth returns true if the given password is correct for the given
// username, and false otherwise.
func (s *Server) VerifyAuth(ctx context.Context, userID types.UserID, password string) (bool, error) {
	p, err := s.db.RetrievePassword(ctx, userID)
	if err != nil {
		return false, err
	}
	return password == p, nil
}
