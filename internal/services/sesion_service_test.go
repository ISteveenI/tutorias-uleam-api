package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
)

type mockSesionRepository struct {
	createCalled bool
}

func (m *mockSesionRepository) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	m.createCalled = true
	return nil
}

func (m *mockSesionRepository) FindByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	return nil, nil
}

func (m *mockSesionRepository) FindAll(ctx context.Context) ([]models.SesionTutoria, error) {
	return nil, nil
}

func (m *mockSesionRepository) Update(ctx context.Context, sesion *models.SesionTutoria) error {
	return nil
}

func (m *mockSesionRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (m *mockSesionRepository) CreateAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	return nil
}

func (m *mockSesionRepository) CreateEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	return nil
}

func TestCreateSesionConSolicitudIDInvalidoNoLlegaAlRepositorio(t *testing.T) {
	mockRepo := &mockSesionRepository{}
	service := services.NewSesionService(mockRepo)

	sesion := &models.SesionTutoria{
		SolicitudID: 0,
		FechaSesion: "2026-07-01",
		HoraInicio:  "09:00",
		HoraFin:     "10:00",
		Estado:      "Programada",
	}

	err := service.Create(context.Background(), sesion)

	if !errors.Is(err, services.ErrDatosInvalidos) {
		t.Fatalf("se esperaba ErrDatosInvalidos, pero se obtuvo: %v", err)
	}

	if mockRepo.createCalled {
		t.Fatal("el repositorio no debio ser llamado cuando la sesion tiene datos invalidos")
	}
}
