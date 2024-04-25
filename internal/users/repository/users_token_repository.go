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

// GetClaims implements users.UsersTokenRepo.
func (u *usersTokenRepository) GetClaims(ctx context.Context, token string) (map[string]interface{}, error) {
	log.Println("GetClaims")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.config.TokenConfig.Secret), nil
	})

	if err != nil {
		log.Printf("Error getting claims: %v", err)
		return nil, err
	}

	log.Println("Claims gotten")
	return claims, nil
}

// CreateToken implements users.UsersTokenRepo.
func (u *usersTokenRepository) CreateToken(ctx context.Context, email string) (string, error) {
	log.Println("CreateToken")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
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
		return []byte(u.config.TokenConfig.Secret), nil
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
