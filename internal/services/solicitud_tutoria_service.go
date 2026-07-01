package services

import (
	"context"
	"strings"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
)

type SolicitudTutoriaServiceInterface interface {
	Create(ctx context.Context, solicitud *models.SolicitudTutoria) error
	GetByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error)
	GetAll(ctx context.Context) ([]models.SolicitudTutoria, error)
	Update(ctx context.Context, id uint, solicitud *models.SolicitudTutoria) error
	Delete(ctx context.Context, id uint) error
}

type SolicitudTutoriaService struct {
	repo repositories.SolicitudTutoriaRepository
}

func NewSolicitudTutoriaService(repo repositories.SolicitudTutoriaRepository) *SolicitudTutoriaService {
	return &SolicitudTutoriaService{
		repo: repo,
	}
}

func (s *SolicitudTutoriaService) Create(ctx context.Context, solicitud *models.SolicitudTutoria) error {

	if err := validarSolicitud(solicitud); err != nil {
		return err
	}

	if strings.TrimSpace(solicitud.Estado) == "" {
		solicitud.Estado = "Pendiente"
	}

	return s.repo.Create(ctx, solicitud)
}

func (s *SolicitudTutoriaService) GetByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {

	if id == 0 {
		return nil, ErrDatosInvalidos
	}

	return s.repo.FindByID(ctx, id)
}

func (s *SolicitudTutoriaService) GetAll(ctx context.Context) ([]models.SolicitudTutoria, error) {
	return s.repo.FindAll(ctx)
}

func (s *SolicitudTutoriaService) Update(ctx context.Context, id uint, solicitud *models.SolicitudTutoria) error {

	if id == 0 {
		return ErrDatosInvalidos
	}

	if err := validarSolicitud(solicitud); err != nil {
		return err
	}

	solicitud.ID = uint(id)

	return s.repo.Update(ctx, solicitud)
}

func (s *SolicitudTutoriaService) Delete(ctx context.Context, id uint) error {

	if id == 0 {
		return ErrDatosInvalidos
	}

	return s.repo.Delete(ctx, id)
}

func validarSolicitud(solicitud *models.SolicitudTutoria) error {

	if solicitud == nil {
		return ErrDatosInvalidos
	}

	if solicitud.EstudianteID <= 0 {
		return ErrDatosInvalidos
	}

	if solicitud.HorarioDocenteID <= 0 {
		return ErrDatosInvalidos
	}

	if solicitud.TipoTutoriaID <= 0 {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(solicitud.Tema) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(solicitud.FechaSolicitud) == "" {
		return ErrDatosInvalidos
	}

	return nil
}