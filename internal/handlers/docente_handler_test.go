package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type FakeDocenteService struct{}

func (f *FakeDocenteService) Create(ctx context.Context, docente *models.Docente) error {
	docente.ID = 1
	return nil
}

func (f *FakeDocenteService) GetByID(ctx context.Context, id uint) (*models.Docente, error) {
	return &models.Docente{
		ID:           id,
		Nombres:      "Karen",
		Apellidos:    "Holguin",
		Correo:       "karen@uleam.edu.ec",
		Departamento: "Tecnologías",
		Especialidad: "Aplicaciones Web",
	}, nil
}

func (f *FakeDocenteService) GetAll(ctx context.Context) ([]models.Docente, error) {
	return []models.Docente{}, nil
}

func (f *FakeDocenteService) Update(ctx context.Context, id uint, docente *models.Docente) error {
	return nil
}

func (f *FakeDocenteService) Delete(ctx context.Context, id uint) error {
	return nil
}

func TestCreateDocente401(t *testing.T) {

	handler := NewDocenteHandler(&FakeDocenteService{})

	body := models.Docente{
		Nombres:      "Karen",
		Apellidos:    "Holguin",
		Correo:       "karen@uleam.edu.ec",
		Departamento: "Tecnologías",
		Especialidad: "Aplicaciones Web",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/",
		bytes.NewBuffer(jsonBody),
	)

	rec := httptest.NewRecorder()

	handler.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf(
			"se esperaba %d y se obtuvo %d",
			http.StatusUnauthorized,
			rec.Code,
		)
	}
}

func TestCreateDocente201(t *testing.T) {

	handler := NewDocenteHandler(&FakeDocenteService{})

	body := models.Docente{
		Nombres:      "Karen",
		Apellidos:    "Holguin",
		Correo:       "karen@uleam.edu.ec",
		Departamento: "Tecnologías",
		Especialidad: "Aplicaciones Web",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Authorization", "Bearer prueba")

	rec := httptest.NewRecorder()

	handler.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusTeapot {
		t.Fatalf(
			"se esperaba %d y se obtuvo %d",
			http.StatusTeapot,
			rec.Code,
		)
	}
}