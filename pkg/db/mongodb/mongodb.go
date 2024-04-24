package mongodb

import (
	"context"
	"fmt"
	"log"
	"url-shortener/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(config *config.Config) (*mongo.Database, error) {
	log.Println("Connecting to MongoDB")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.MongoConfig.Host, config.MongoConfig.Port)))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	clientDb := client.Database(config.MongoConfig.Database)
	log.Println("Connected to MongoDB")
	return clientDb, nil
}
