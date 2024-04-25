package url

import (
	"context"
	"time"
)

type UrlRedisRepository interface {
	SaveUrl(ctx context.Context, urlToken string, urlValue string, durantion time.Duration) error
	GetUrl(ctx context.Context, urlToken string) (string, error)
}
