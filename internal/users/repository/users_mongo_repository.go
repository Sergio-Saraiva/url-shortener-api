package repository

import (
	"context"
	"log"
	"url-shortener/internal/models"
	"url-shortener/internal/users"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersMongoRepository struct {
	mongoClient *mongo.Collection
}

func NewUsersMongoRepository(mongoClient *mongo.Database) users.UsersMongoRepo {
	return &usersMongoRepository{
		mongoClient: mongoClient.Collection("users"),
	}
}

// CreateUser implements users.UsersMongoRepo.
func (u *usersMongoRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("Creating user")
	user.ID = uuid.NewString()
	_, err := u.mongoClient.InsertOne(ctx, &user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	log.Println("User created")
	return user, nil
}

// DeleteUser implements users.UsersMongoRepo.
func (u *usersMongoRepository) DeleteUser(ctx context.Context, user *models.User) error {
	log.Println("Deleting user")
	result, err := u.mongoClient.DeleteOne(ctx, user)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		log.Println("User not found")
		return nil
	}

	log.Println("User deleted")
	return nil
}

// GetUser implements users.UsersMongoRepo.
func (u *usersMongoRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	log.Println("Getting user")
	var user models.User
	filter := bson.D{{Key: "email", Value: email}}
	err := u.mongoClient.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("User not found")
			return nil, nil
		}

		log.Printf("Error getting user: %v", err)
		return nil, err
	}

	log.Println("User retrieved")
	return &user, nil
}
