package usecases

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"url-shortener/config"
	"url-shortener/internal/url"
)

type urlUseCase struct {
	redisRepo url.UrlRedisRepository
	cfg       *config.Config
}

// GetUrl implements url.UrlUseCase.
func (u *urlUseCase) GetUrl(ctx context.Context, urlToken string) (string, error) {
	log.Println("Getting URL from Redis")
	result, err := u.redisRepo.GetUrl(ctx, urlToken)
	if err != nil {
		log.Printf("Error getting URL from Redis: %v", err)
		return "", err
	}

	log.Println("URL retrieved from Redis")
	return result, nil
}

func NewUrlUseCase(redisRepo url.UrlRedisRepository, cfg *config.Config) url.UrlUseCase {
	return &urlUseCase{
		redisRepo: redisRepo,
		cfg:       cfg,
	}
}

// GenerateShortUrl implements url.UrlUseCase.
func (u *urlUseCase) GenerateShortUrl(ctx context.Context, urlToken string) string {
	result := fmt.Sprintf("http://%s:%d/r/%s", u.cfg.ServerConfig.Host, u.cfg.ServerConfig.Port, urlToken)
	return result
}

// GenerateUrlToken implements url.UrlUseCase.
func (u *urlUseCase) GenerateUrlToken(ctx context.Context, url string) string {
	log.Println("Generating URL token")
	b := make([]byte, 6)
	for i := range b {
		b[i] = url[rand.Intn(len(url))]
	}
	log.Println("URL token generated")
	return string(b)
}

// SaveUrl implements url.UrlUseCase.
func (u *urlUseCase) SaveUrl(ctx context.Context, urlToken string, urlValue string) error {
	log.Println("Saving URL to Redis")
	err := u.redisRepo.SaveUrl(ctx, urlToken, urlValue)
	if err != nil {
		log.Printf("Error saving URL to Redis: %v", err)
		return err
	}

	log.Println("URL saved to Redis")
	return nil
}
