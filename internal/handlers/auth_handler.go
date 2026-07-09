package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	r.Group(func(protegida chi.Router) {
		protegida.Use(middlewares.RequireAuth)
		protegida.Get("/perfil", h.Perfil)
	})

	r.Group(func(admin chi.Router) {
		admin.Use(middlewares.RequireRole("admin"))
		admin.Get("/admin", h.Admin)
	})

	return r
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req services.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		escribirAuthJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON invalido",
		})
		return
	}

	usuario, err := h.service.Register(r.Context(), req)
	if err != nil {
		manejarErrorAuth(w, err)
		return
	}

	escribirAuthJSON(w, http.StatusCreated, usuario)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req services.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		escribirAuthJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON invalido",
		})
		return
	}

	resp, err := h.service.Login(r.Context(), req)
	if err != nil {
		manejarErrorAuth(w, err)
		return
	}

	escribirAuthJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) Perfil(w http.ResponseWriter, r *http.Request) {
	escribirAuthJSON(w, http.StatusOK, map[string]any{
		"usuario_id": r.Context().Value(middlewares.UsuarioIDKey),
		"correo":     r.Context().Value(middlewares.CorreoKey),
		"rol":        r.Context().Value(middlewares.RolKey),
	})
}

func (h *AuthHandler) Admin(w http.ResponseWriter, r *http.Request) {
	escribirAuthJSON(w, http.StatusOK, map[string]string{
		"mensaje": "Acceso permitido solo para rol admin",
	})
}

func manejarErrorAuth(w http.ResponseWriter, err error) {
	if errors.Is(err, services.ErrDatosInvalidos) {
		escribirAuthJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Datos invalidos",
		})
		return
	}

	if errors.Is(err, services.ErrRolInvalido) {
		escribirAuthJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Rol invalido",
		})
		return
	}

	if errors.Is(err, services.ErrCredencialesInvalidas) {
		escribirAuthJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "Credenciales invalidas",
		})
		return
	}

	escribirAuthJSON(w, http.StatusInternalServerError, map[string]string{
		"error": "Error interno del servidor",
	})
}

func escribirAuthJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
