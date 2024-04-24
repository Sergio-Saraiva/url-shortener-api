package users

import (
	"context"
	"url-shortener/internal/models"
)

type UsersUseCases interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	SignIn(ctx context.Context, user *models.User, signInReq *models.SignInRequest) (string, error)
	DeleteUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, email string) (*models.User, error)
}
