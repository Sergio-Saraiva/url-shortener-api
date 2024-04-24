package users

import (
	"context"
	"url-shortener/internal/models"
)

type UsersMongoRepo interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, email string) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) error
}
