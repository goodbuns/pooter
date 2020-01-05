package api

import "context"

// VerifyAuth returns true if the given password is correct for the given
// username, and false otherwise.
func (s *Server) VerifyAuth(ctx context.Context, userID, password string) (bool, error) {
	p, err := s.DB.RetrievePassword(ctx, userID)
	if err != nil {
		return false, err
	}
	return password == p, nil
}
