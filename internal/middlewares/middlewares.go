package middlewares

import "url-shortener/internal/users"

type MiddlewaresManager struct {
	tokenRepo   users.UsersTokenRepo
	userUseCase users.UsersUseCases
}

func NewMiddlewaresManager(tokenRepo users.UsersTokenRepo, userUseCase users.UsersUseCases) *MiddlewaresManager {
	return &MiddlewaresManager{
		tokenRepo:   tokenRepo,
		userUseCase: userUseCase,
	}
}
