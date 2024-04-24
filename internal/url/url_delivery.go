package url

import "net/http"

type UrlHandler interface {
	CreateShortUrl() http.HandlerFunc
	RedirectToOriginalUrl() http.HandlerFunc
}
