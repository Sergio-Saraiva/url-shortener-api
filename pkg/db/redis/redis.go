package redis

import (
	"context"
	"log"
	"url-shortener/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	log.Println("Creating redis client")
	redisHost := cfg.RedisConfig.RedisAddr

	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	})

	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
		return nil
	}

	log.Printf("Connected to redis: %v", result)
	return client
}
