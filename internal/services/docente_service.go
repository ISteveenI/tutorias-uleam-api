package services

import (
	"context"
	//"errors"
	"strings"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
)

//var ErrDatosInvalidos = errors.New("datos invalidos")

type DocenteServiceInterface interface {
	Create(ctx context.Context, docente *models.Docente) error
	GetByID(ctx context.Context, id uint) (*models.Docente, error)
	GetAll(ctx context.Context) ([]models.Docente, error)
	Update(ctx context.Context, id uint, docente *models.Docente) error
	Delete(ctx context.Context, id uint) error
}

type DocenteService struct {
	repo repositories.DocenteRepository
}

func NewDocenteService(repo repositories.DocenteRepository) *DocenteService {
	return &DocenteService{
		repo: repo,
	}
}

func (s *DocenteService) Create(
	ctx context.Context,
	docente *models.Docente,
) error {

	if err := validarDocente(docente); err != nil {
		return err
	}

	return s.repo.Create(ctx, docente)
}

func (s *DocenteService) GetByID(
	ctx context.Context,
	id uint,
) (*models.Docente, error) {

	if id == 0 {
		return nil, ErrDatosInvalidos
	}

	return s.repo.FindByID(ctx, id)
}

func (s *DocenteService) GetAll(
	ctx context.Context,
) ([]models.Docente, error) {

	return s.repo.FindAll(ctx)
}

func (s *DocenteService) Update(
	ctx context.Context,
	id uint,
	docente *models.Docente,
) error {

	if id == 0 {
		return ErrDatosInvalidos
	}

	if err := validarDocente(docente); err != nil {
		return err
	}

	docente.ID = id

	return s.repo.Update(ctx, docente)
}

func (s *DocenteService) Delete(
	ctx context.Context,
	id uint,
) error {

	if id == 0 {
		return ErrDatosInvalidos
	}

	return s.repo.Delete(ctx, id)
}

func validarDocente(docente *models.Docente) error {

	if docente == nil {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(docente.Nombres) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(docente.Apellidos) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(docente.Correo) == "" {
		return ErrDatosInvalidos
	}

	if !strings.Contains(docente.Correo, "@") {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(docente.Departamento) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(docente.Especialidad) == "" {
		return ErrDatosInvalidos
	}

	return nil
}