package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/auth"
)

func RequireRole(roles ...string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			value := r.Context().Value(UserContextKey)

			claims, ok := value.(*auth.Claims)

			if !ok {

				w.WriteHeader(http.StatusUnauthorized)

				json.NewEncoder(w).Encode(map[string]string{
					"error": "usuario no autenticado",
				})

				return
			}

			for _, role := range roles {

				if claims.Role == role {

					next.ServeHTTP(w, r)
					return
				}
			}

			w.WriteHeader(http.StatusForbidden)

			json.NewEncoder(w).Encode(map[string]string{
				"error": "sin permisos",
			})
		})
	}
}
