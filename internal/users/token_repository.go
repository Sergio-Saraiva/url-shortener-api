package users

import "context"

type UsersTokenRepo interface {
	CreateToken(ctx context.Context, email string) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
	GetClaims(ctx context.Context, token string) (map[string]interface{}, error)
}
