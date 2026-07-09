package services

import (
	"context"
	"errors"
	"strings"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var ErrCredencialesInvalidas = errors.New("credenciales invalidas")
var ErrRolInvalido = errors.New("rol invalido")

type RegisterRequest struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}

type LoginRequest struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token   string         `json:"token"`
	Usuario models.Usuario `json:"usuario"`
}

type AuthService struct {
	repo repositories.UsuarioRepository
}

func NewAuthService(repo repositories.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (*models.Usuario, error) {
	req.Nombre = strings.TrimSpace(req.Nombre)
	req.Correo = strings.TrimSpace(req.Correo)
	req.Rol = strings.TrimSpace(req.Rol)

	if req.Rol == "" {
		req.Rol = "estudiante"
	}

	if req.Nombre == "" || req.Correo == "" || strings.TrimSpace(req.Password) == "" {
		return nil, ErrDatosInvalidos
	}

	if !rolValido(req.Rol) {
		return nil, ErrRolInvalido
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	usuario := &models.Usuario{
		Nombre:       req.Nombre,
		Correo:       req.Correo,
		PasswordHash: string(hash),
		Rol:          req.Rol,
	}

	if err := s.repo.Create(ctx, usuario); err != nil {
		return nil, err
	}

	return usuario, nil
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	req.Correo = strings.TrimSpace(req.Correo)

	if req.Correo == "" || strings.TrimSpace(req.Password) == "" {
		return nil, ErrCredencialesInvalidas
	}

	usuario, err := s.repo.FindByCorreo(ctx, req.Correo)
	if err != nil {
		return nil, ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrCredencialesInvalidas
	}

	token, err := middlewares.GenerateToken(usuario.ID, usuario.Correo, usuario.Rol)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:   token,
		Usuario: *usuario,
	}, nil
}

func rolValido(rol string) bool {
	switch rol {
	case "admin", "docente", "estudiante":
		return true
	default:
		return false
	}
}
