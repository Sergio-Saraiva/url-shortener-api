package usecases

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/users"

	"github.com/goombaio/namegenerator"
)

type usersUseCases struct {
	usersRepo     users.UsersMongoRepo
	userTokenRepo users.UsersTokenRepo
}

func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}

// CreateUser implements users.UsersUseCases.
func (u *usersUseCases) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("Creating user usecase")

	user.Password = encryptPassword(user.Password)
	user.Username = namegenerator.NewNameGenerator(time.Now().UTC().UnixNano()).Generate()
	user.Type = "paid"

	createdUser, err := u.usersRepo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	log.Println("User created usecase")
	return createdUser, nil
}

// DeleteUser implements users.UsersUseCases.
func (u *usersUseCases) DeleteUser(ctx context.Context, user *models.User) error {
	log.Println("Deleting user usecase")
	err := u.usersRepo.DeleteUser(ctx, user)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	log.Println("User deleted usecase")
	return nil
}

// GetUser implements users.UsersUseCases.
func (u *usersUseCases) GetUser(ctx context.Context, email string) (*models.User, error) {
	log.Println("Getting user usecase")
	user, err := u.usersRepo.GetUser(ctx, email)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}

	log.Println("User gotten usecase")
	return user, nil
}

// SignIn implements users.UsersUseCases.
func (u *usersUseCases) SignIn(ctx context.Context, user *models.User, signInReq *models.SignInRequest) (string, error) {
	log.Println("Signing in usecase")

	signInReq.Password = encryptPassword(signInReq.Password)

	if signInReq.Password != user.Password {
		log.Println("Invalid password")
		return "", errors.New("invalid password")
	}

	token, err := u.userTokenRepo.CreateToken(ctx, user.Email)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return "", err
	}

	log.Println("Token created usecase")
	return token, nil
}

func NewUsersUseCases(usersRepo users.UsersMongoRepo, usersTokenRepo users.UsersTokenRepo) users.UsersUseCases {
	return &usersUseCases{
		usersRepo:     usersRepo,
		userTokenRepo: usersTokenRepo,
	}
}
