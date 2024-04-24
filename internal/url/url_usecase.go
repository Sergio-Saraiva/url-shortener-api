package url

import "context"

type UrlUseCase interface {
	SaveUrl(ctx context.Context, urlToken string, urlValue string) error
	GenerateShortUrl(ctx context.Context, url string) string
	GenerateUrlToken(ctx context.Context, url string) string
	GetUrl(ctx context.Context, urlToken string) (string, error)
	GenerateQRCode(ctx context.Context, url string) ([]byte, error)
}
