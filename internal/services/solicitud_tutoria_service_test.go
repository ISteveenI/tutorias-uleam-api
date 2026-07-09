package services

import (
	"context"
	"errors"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type mockSolicitudRepository struct {
	createCalled   bool
	updateCalled   bool
	deleteCalled   bool
	findAllCalled  bool
	findByIDCalled bool
}

func (m *mockSolicitudRepository) Create(ctx context.Context, s *models.SolicitudTutoria) error {
	m.createCalled = true
	return nil
}

func (m *mockSolicitudRepository) FindByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {
	m.findByIDCalled = true

	if id == 999 {
		return nil, errors.New("no encontrado")
	}

	return &models.SolicitudTutoria{
		ID:               id,
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
		Estado:           "Pendiente",
	}, nil
}

func (m *mockSolicitudRepository) FindAll(ctx context.Context) ([]models.SolicitudTutoria, error) {
	m.findAllCalled = true

	return []models.SolicitudTutoria{
		{
			ID:               1,
			EstudianteID:     1,
			HorarioDocenteID: 1,
			TipoTutoriaID:    1,
			Tema:             "POO",
			FechaSolicitud:   "2026-06-30",
			Estado:           "Pendiente",
		},
	}, nil
}

func (m *mockSolicitudRepository) Update(ctx context.Context, s *models.SolicitudTutoria) error {
	m.updateCalled = true
	return nil
}

func (m *mockSolicitudRepository) Delete(ctx context.Context, id uint) error {
	m.deleteCalled = true
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

func TestCreateSolicitud_EstadoPendientePorDefecto(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "Programación",
		FechaSolicitud:   "2026-06-30",
	}

	err := service.Create(context.Background(), solicitud)

	if err != nil {
		t.Fatal(err)
	}

	if !mockRepo.createCalled {
		t.Fatal("el repository debía ejecutarse")
	}

	if solicitud.Estado != "Pendiente" {
		t.Fatal("el estado debía ser Pendiente")
	}
}

func TestGetAll(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	lista, err := service.GetAll(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if !mockRepo.findAllCalled {
		t.Fatal("debía llamar al repository")
	}

	if len(lista) != 1 {
		t.Fatal("se esperaba una solicitud")
	}
}

func TestGetByID_OK(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	solicitud, err := service.GetByID(context.Background(), 1)

	if err != nil {
		t.Fatal(err)
	}

	if !mockRepo.findByIDCalled {
		t.Fatal("debía llamar al repository")
	}

	if solicitud.ID != 1 {
		t.Fatal("id incorrecto")
	}
}

func TestGetByID_IDInvalido(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	_, err := service.GetByID(context.Background(), 0)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}

func TestUpdate_OK(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
	}

	err := service.Update(context.Background(), 1, solicitud)

	if err != nil {
		t.Fatal(err)
	}

	if !mockRepo.updateCalled {
		t.Fatal("debía llamar al repository")
	}
}

func TestUpdate_IDInvalido(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
	}

	err := service.Update(context.Background(), 0, solicitud)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}

func TestDelete_OK(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	err := service.Delete(context.Background(), 1)

	if err != nil {
		t.Fatal(err)
	}

	if !mockRepo.deleteCalled {
		t.Fatal("debía llamar al repository")
	}
}

func TestDelete_IDInvalido(t *testing.T) {

	mockRepo := &mockSolicitudRepository{}
	service := NewSolicitudTutoriaService(mockRepo)

	err := service.Delete(context.Background(), 0)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}
