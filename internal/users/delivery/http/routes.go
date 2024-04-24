package http

import (
	"url-shortener/internal/users"

	"github.com/go-chi/chi/v5"
)

func MapUsersRoutes(router *chi.Mux, handler users.UsersDelivery) {
	router.Post("/signup", handler.CreateUser())
	router.Post("/signin", handler.SignIn())
}
