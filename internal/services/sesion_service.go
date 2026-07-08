package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
)

var ErrDatosInvalidos = errors.New("datos invalidos")

type SesionServiceInterface interface {
	Create(ctx context.Context, sesion *models.SesionTutoria) error
	GetByID(ctx context.Context, id uint) (*models.SesionTutoria, error)
	GetAll(ctx context.Context) ([]models.SesionTutoria, error)
	Update(ctx context.Context, id uint, sesion *models.SesionTutoria) error
	Delete(ctx context.Context, id uint) error

	RegistrarAsistencia(ctx context.Context, asistencia *models.Asistencia) error
	SubirEvidencia(ctx context.Context, evidencia *models.Evidencia) error
}

type SesionService struct {
	repo repositories.SesionRepository

	// estadoPorDefecto permite configurar el estado inicial de una sesión
	// sin dejar el valor quemado directamente dentro del método Create.
	estadoPorDefecto string
}

// SesionOption representa una configuración opcional para SesionService.
// Se usa el patrón Functional Options para no modificar la forma tradicional
// de crear el service cuando no se necesita configuración extra.
type SesionOption func(*SesionService)

// WithEstadoPorDefecto permite cambiar el estado inicial usado cuando
// una sesión se crea sin estado explícito.
func WithEstadoPorDefecto(estado string) SesionOption {
	return func(s *SesionService) {
		if strings.TrimSpace(estado) != "" {
			s.estadoPorDefecto = estado
		}
	}
}

// NewSesionService crea el service de sesiones.
// Por defecto usa el estado "Programada", pero permite recibir opciones
// para personalizar valores sin romper las llamadas existentes.
func NewSesionService(repo repositories.SesionRepository, opts ...SesionOption) *SesionService {
	service := &SesionService{
		repo:             repo,
		estadoPorDefecto: "Programada",
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

func (s *SesionService) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	if err := validarSesion(sesion); err != nil {
		return err
	}
	// Si el cliente no envía un estado, se usa el estado configurado
	// en el service mediante Functional Options.
	if strings.TrimSpace(sesion.Estado) == "" {
		sesion.Estado = s.estadoPorDefecto
	}

	return s.repo.Create(ctx, sesion)
}

func (s *SesionService) GetByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	if id == 0 {
		return nil, ErrDatosInvalidos
	}

	return s.repo.FindByID(ctx, id)
}

func (s *SesionService) GetAll(ctx context.Context) ([]models.SesionTutoria, error) {
	return s.repo.FindAll(ctx)
}

func (s *SesionService) Update(ctx context.Context, id uint, sesion *models.SesionTutoria) error {
	if id == 0 {
		return ErrDatosInvalidos
	}

	if err := validarSesion(sesion); err != nil {
		return err
	}

	sesion.ID = id
	return s.repo.Update(ctx, sesion)
}

func (s *SesionService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return ErrDatosInvalidos
	}

	return s.repo.Delete(ctx, id)
}

func (s *SesionService) RegistrarAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	if asistencia.SesionID == 0 {
		return ErrDatosInvalidos
	}

	return s.repo.CreateAsistencia(ctx, asistencia)
}

func (s *SesionService) SubirEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	if evidencia.SesionID == 0 {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(evidencia.TipoArchivo) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(evidencia.ArchivoURL) == "" {
		return ErrDatosInvalidos
	}

	if evidencia.FechaSubida.IsZero() {
		evidencia.FechaSubida = time.Now()
	}

	return s.repo.CreateEvidencia(ctx, evidencia)
}

func validarSesion(sesion *models.SesionTutoria) error {
	if sesion == nil {
		return ErrDatosInvalidos
	}

	if sesion.SolicitudID == 0 {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(sesion.FechaSesion) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(sesion.HoraInicio) == "" {
		return ErrDatosInvalidos
	}

	if strings.TrimSpace(sesion.HoraFin) == "" {
		return ErrDatosInvalidos
	}

	if sesion.HoraInicio >= sesion.HoraFin {
		return ErrDatosInvalidos
	}

	return nil
}
