package http

import (
	"url-shortener/internal/middlewares"
	"url-shortener/internal/url"

	"github.com/go-chi/chi/v5"
)

func MapUrlShortenerRoutes(router *chi.Mux, handler url.UrlHandler, mwManager middlewares.MiddlewaresManager) {
	router.Get("/r/{urlToken}", handler.RedirectToOriginalUrl())
	router.Post("/qrcode", handler.CreateQRCode())
	router.Group(func(router chi.Router) {
		router.Use(mwManager.AuthMiddleware)
		router.Post("/shorten", handler.CreateShortUrl())
	})
}
