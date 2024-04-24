package server

import (
	"net/http"
	"url-shortener/internal/url/repository"
	"url-shortener/internal/url/usecases"

	UrlHandler "url-shortener/internal/url/delivery/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) MapHandlers(chi *chi.Mux) error {
	chi.Use(middleware.Logger)

	urlRedisRepo := repository.NewUrlRedisRepository(s.redisClient)
	urlUseCases := usecases.NewUrlUseCase(urlRedisRepo, s.config)
	urlHandler := UrlHandler.NewAuthHandlers(s.config, urlUseCases)
	UrlHandler.MapUrlShortenerRoutes(chi, urlHandler)

	chi.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	})

	return nil
}
