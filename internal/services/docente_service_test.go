package services

import (
	"context"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
)

type MockDocenteRepository struct {
	createCalled bool
}

func (m *MockDocenteRepository) Create(ctx context.Context, docente *models.Docente) error {
	m.createCalled = true
	return nil
}

func (m *MockDocenteRepository) FindByID(ctx context.Context, id uint) (*models.Docente, error) {
	return nil, nil
}

func (m *MockDocenteRepository) FindAll(ctx context.Context) ([]models.Docente, error) {
	return []models.Docente{}, nil
}

func (m *MockDocenteRepository) Update(ctx context.Context, docente *models.Docente) error {
	return nil
}

func (m *MockDocenteRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

var _ repositories.DocenteRepository = (*MockDocenteRepository)(nil)

func TestDocenteService_NoDebeCrearDocenteSinCorreo(t *testing.T) {

	mockRepo := &MockDocenteRepository{}

	service := NewDocenteService(mockRepo)

	docente := &models.Docente{
		Nombres:   "Karen",
		Apellidos: "Holguin",
		Correo:    "",
	}

	err := service.Create(context.Background(), docente)

	if err == nil {
		t.Fatal("se esperaba un error")
	}

	if mockRepo.createCalled {
		t.Fatal("no debía llamar al repositorio")
	}
}
