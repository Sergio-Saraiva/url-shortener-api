package http

import (
	"url-shortener/internal/url"

	"github.com/go-chi/chi/v5"
)

func MapUrlShortenerRoutes(router *chi.Mux, handler url.UrlHandler) {
	router.Post("/shorten", handler.CreateShortUrl())
	router.Get("/r/{urlToken}", handler.RedirectToOriginalUrl())
	router.Post("/qrcode", handler.CreateQRCode())
}
