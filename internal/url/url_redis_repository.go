package url

import "context"

type UrlRedisRepository interface {
	SaveUrl(ctx context.Context, urlToken string, urlValue string) error
	GetUrl(ctx context.Context, urlToken string) (string, error)
}
