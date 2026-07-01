package services

import (
	"context"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type mockSolicitudRepository struct {
	createCalled bool
}

func (m *mockSolicitudRepository) Create(ctx context.Context, s *models.SolicitudTutoria) error {
	m.createCalled = true
	return nil
}

func (m *mockSolicitudRepository) FindByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {
	return nil, nil
}

func (m *mockSolicitudRepository) FindAll(ctx context.Context) ([]models.SolicitudTutoria, error) {
	return nil, nil
}

func (m *mockSolicitudRepository) Update(ctx context.Context, s *models.SolicitudTutoria) error {
	return nil
}

func (m *mockSolicitudRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func TestCreateSolicitud_DatosInvalidos_NoLlamaRepository(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}

	service := NewSolicitudTutoriaService(mockRepo)

	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "",
		FechaSolicitud:   "2026-06-30",
	}

	err := service.Create(context.Background(), solicitud)

	if err == nil {
		t.Fatal("se esperaba un error")
	}

	if mockRepo.createCalled {
		t.Fatal("el repository NO debía ejecutarse")
	}
}
