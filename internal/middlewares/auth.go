package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"
)

type ctxUserKey string

const (
	UserIDKey ctxUserKey = "c-user-id"
	TypeKey   ctxUserKey = "c-type"
)

func (m *MiddlewaresManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ctx := context.WithValue(r.Context(), UserIDKey, nil)
			ctx = context.WithValue(ctx, TypeKey, "unpaid")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		token := strings.Split(authHeader, " ")[1]
		_, err := m.tokenRepo.ValidateToken(r.Context(), token)
		if err != nil {
			log.Printf("Error validating token: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, err := m.tokenRepo.GetClaims(r.Context(), token)
		if err != nil {
			log.Printf("Error getting claims: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
		}

		user, err := m.userUseCase.GetUser(r.Context(), claims["email"].(string))
		if err != nil {
			log.Printf("Error getting user: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
		ctx = context.WithValue(ctx, TypeKey, user.Type)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
