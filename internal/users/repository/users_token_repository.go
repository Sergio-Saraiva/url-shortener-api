package repository

import (
	"context"
	"fmt"
	"log"
	"time"
	"url-shortener/config"
	"url-shortener/internal/users"

	"github.com/golang-jwt/jwt/v5"
)

type usersTokenRepository struct {
	config *config.Config
}

// CreateToken implements users.UsersTokenRepo.
func (u *usersTokenRepository) CreateToken(ctx context.Context, username string) (string, error) {
	log.Println("CreateToken")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenBytes, err := token.SignedString([]byte(u.config.TokenConfig.Secret))
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return "", err
	}

	log.Println("Token created")
	return tokenBytes, nil
}

// ValidateToken implements users.UsersTokenRepo.
func (u *usersTokenRepository) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	log.Println("ValidateToken")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return u.config.TokenConfig.Secret, nil
	})

	if err != nil {
		log.Printf("Error validating token: %v", err)
		return false, err
	}

	if !token.Valid {
		log.Println("Invalid token")
		return false, fmt.Errorf("invalid token")
	}

	log.Println("Token validated")
	return true, nil
}

func NewUsersTokenRepository(config *config.Config) users.UsersTokenRepo {
	return &usersTokenRepository{config: config}
}
