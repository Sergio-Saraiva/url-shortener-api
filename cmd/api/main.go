package main

import (
	"log"
	"os"
	"url-shortener/config"
	"url-shortener/internal/server"
	"url-shortener/pkg/db/mongodb"
	"url-shortener/pkg/db/redis"
	"url-shortener/pkg/utils"
)

func main() {
	log.Println("Starting API server...")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	redisClient := redis.NewRedisClient(cfg)

	mongoClient, err := mongodb.NewMongoDBConnection(cfg)
	if err != nil {
		log.Fatalf("Error creating connection to MongoDB: %v", err)
	}

	s := server.NewServer(cfg, redisClient, mongoClient)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
