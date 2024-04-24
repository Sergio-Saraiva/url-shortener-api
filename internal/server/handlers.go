package server

import (
	"net/http"
	urlRedisRepo "url-shortener/internal/url/repository"
	ursUseCase "url-shortener/internal/url/usecases"
	usersDeliveryHttp "url-shortener/internal/users/delivery/http"
	usersRepo "url-shortener/internal/users/repository"
	usersUsecases "url-shortener/internal/users/usecases"

	urlDeliveryHttp "url-shortener/internal/url/delivery/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) MapHandlers(chi *chi.Mux) error {
	chi.Use(middleware.Logger)

	urlRedisRepo := urlRedisRepo.NewUrlRedisRepository(s.redisClient)
	usersMongorepo := usersRepo.NewUsersMongoRepository(s.mongoClient)
	usersTokenRepo := usersRepo.NewUsersTokenRepository(s.config)

	urlUseCases := ursUseCase.NewUrlUseCase(urlRedisRepo, s.config)
	usersUseCases := usersUsecases.NewUsersUseCases(usersMongorepo, usersTokenRepo)

	urlHandler := urlDeliveryHttp.NewAuthHandlers(s.config, urlUseCases)
	urlDeliveryHttp.MapUrlShortenerRoutes(chi, urlHandler)

	usersHandler := usersDeliveryHttp.NewHttpUserHandler(usersUseCases)
	usersDeliveryHttp.MapUsersRoutes(chi, usersHandler)

	chi.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	})

	return nil
}
