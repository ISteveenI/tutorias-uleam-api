package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/auth"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

func RequireAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"error": "token requerido",
			})

			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := auth.ValidateToken(token)

		if err != nil {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"error": "token inválido",
			})

			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserContextKey,
			claims,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
