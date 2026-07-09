package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
)

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {

	if user.Username == "" ||
		user.Password == "" ||
		user.Role == "" {
		return errors.New("todos los campos son obligatorios")
	}

	_, err := s.repo.FindByUsername(ctx, user.Username)

	if err == nil {
		return errors.New("el usuario ya existe")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	return s.repo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*models.User, error) {

	user, err := s.repo.FindByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	return user, nil
}
