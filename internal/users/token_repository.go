package users

import "context"

type UsersTokenRepo interface {
	CreateToken(ctx context.Context, username string) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
}
