package repository

import (
	"context"
	"log"
	"time"
	"url-shortener/internal/url"

	"github.com/redis/go-redis/v9"
)

type urlRedisRepository struct {
	redisClient *redis.Client
}

func NewUrlRedisRepository(redisClient *redis.Client) url.UrlRedisRepository {
	return &urlRedisRepository{redisClient: redisClient}
}

// GetUrl implements url.UrlRedisRepository.
func (u *urlRedisRepository) GetUrl(ctx context.Context, urlToken string) (string, error) {
	log.Println("Getting URL from Redis")
	result, err := u.redisClient.Get(ctx, urlToken).Result()
	if err != nil {
		log.Printf("Error getting URL from Redis: %v", err)
		return "", err
	}

	log.Printf("URL retrieved from Redis: %v", result)
	return result, nil
}

// SaveUrl implements url.UrlRedisRepository.
func (u *urlRedisRepository) SaveUrl(ctx context.Context, urlToken string, urlValue string, duration time.Duration) error {
	log.Println("Saving URL to Redis")
	result, err := u.redisClient.Set(ctx, urlToken, urlValue, duration).Result()
	if err != nil {
		log.Printf("Error saving URL to Redis: %v", err)
		return err
	}

	log.Printf("URL saved to Redis: %v", result)
	return nil
}
