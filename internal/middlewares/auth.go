package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UsuarioIDKey contextKey = "usuario_id"
	CorreoKey    contextKey = "correo"
	RolKey       contextKey = "rol"
)

type AuthClaims struct {
	UsuarioID uint
	Correo    string
	Rol       string
}

func jwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if strings.TrimSpace(secret) == "" {
		secret = "secret-hito3"
	}

	return []byte(secret)
}

func GenerateToken(usuarioID uint, correo string, rol string) (string, error) {
	claims := jwt.MapClaims{
		"usuario_id": usuarioID,
		"correo":     correo,
		"rol":        rol,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret())
}

func obtenerTokenDesdeHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if strings.TrimSpace(authHeader) == "" {
		return "", errors.New("token requerido")
	}

	partes := strings.SplitN(authHeader, " ", 2)
	if len(partes) != 2 || strings.ToLower(partes[0]) != "bearer" {
		return "", errors.New("formato de token invalido")
	}

	return strings.TrimSpace(partes[1]), nil
}

func ParseTokenFromRequest(r *http.Request) (*AuthClaims, error) {
	tokenString, err := obtenerTokenDesdeHeader(r)
	if err != nil {
		return nil, err
	}

	// Token de desarrollo usado por las pruebas antiguas y pruebas rapidas con curl.
	// Para la defensa debe usarse el token generado por /api/v1/auth/login.
	if tokenString == "token-prueba" {
		return &AuthClaims{
			UsuarioID: 1,
			Correo:    "dev@uleam.edu.ec",
			Rol:       "admin",
		}, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metodo de firma invalido")
		}

		return jwtSecret(), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token invalido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims invalidos")
	}

	usuarioIDFloat, ok := claims["usuario_id"].(float64)
	if !ok {
		return nil, errors.New("usuario_id invalido")
	}

	correo, _ := claims["correo"].(string)
	rol, _ := claims["rol"].(string)

	if strings.TrimSpace(correo) == "" || strings.TrimSpace(rol) == "" {
		return nil, errors.New("token incompleto")
	}

	return &AuthClaims{
		UsuarioID: uint(usuarioIDFloat),
		Correo:    correo,
		Rol:       rol,
	}, nil
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := ParseTokenFromRequest(r)
		if err != nil {
			escribirAuthError(w, http.StatusUnauthorized, "token invalido o requerido")
			return
		}

		ctx := context.WithValue(r.Context(), UsuarioIDKey, claims.UsuarioID)
		ctx = context.WithValue(ctx, CorreoKey, claims.Correo)
		ctx = context.WithValue(ctx, RolKey, claims.Rol)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireRole(rolesPermitidos ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := ParseTokenFromRequest(r)
			if err != nil {
				escribirAuthError(w, http.StatusUnauthorized, "token invalido o requerido")
				return
			}

			if !rolPermitido(claims.Rol, rolesPermitidos) {
				escribirAuthError(w, http.StatusForbidden, "rol no autorizado")
				return
			}

			ctx := context.WithValue(r.Context(), UsuarioIDKey, claims.UsuarioID)
			ctx = context.WithValue(ctx, CorreoKey, claims.Correo)
			ctx = context.WithValue(ctx, RolKey, claims.Rol)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func rolPermitido(rol string, rolesPermitidos []string) bool {
	for _, permitido := range rolesPermitidos {
		if rol == permitido {
			return true
		}
	}

	return false
}

func escribirAuthError(w http.ResponseWriter, status int, mensaje string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": mensaje,
	})
}
