package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type fakeSolicitudService struct {
	data []models.SolicitudTutoria
}

func (f *fakeSolicitudService) Create(ctx context.Context, s *models.SolicitudTutoria) error {
	s.ID = uint(len(f.data) + 1)
	f.data = append(f.data, *s)
	return nil
}

func (f *fakeSolicitudService) GetByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {
	for _, v := range f.data {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("solicitud no encontrada")
}

func (f *fakeSolicitudService) GetAll(ctx context.Context) ([]models.SolicitudTutoria, error) {
	return f.data, nil
}

func (f *fakeSolicitudService) Update(ctx context.Context, id uint, s *models.SolicitudTutoria) error {
	return nil
}

func (f *fakeSolicitudService) Delete(ctx context.Context, id uint) error {
	return nil
}

func TestCreateSolicitudHandler(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	body := models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/solicitudes-tutoria",
		bytes.NewReader(jsonBody),
	)

	rec := httptest.NewRecorder()

	handler.CreateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201 y obtuvo %d", rec.Code)
	}

	if len(fake.data) != 1 {
		t.Fatal("el fake debía guardar una solicitud")
	}
}

func TestProtectedRouteUnauthorized(t *testing.T) {

	router := chi.NewRouter()

	router.With(
		middlewares.RequireAuth,
	).Post("/solicitudes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	req := httptest.NewRequest(
		http.MethodPost,
		"/solicitudes",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 y obtuvo %d", rec.Code)
	}
}

func TestGetSolicitudesHandler(t *testing.T) {

	fake := &fakeSolicitudService{
		data: []models.SolicitudTutoria{
			{
				ID:               1,
				EstudianteID:     1,
				HorarioDocenteID: 1,
				TipoTutoriaID:    1,
				Tema:             "POO",
				FechaSolicitud:   "2026-06-30",
				Estado:           "Pendiente",
			},
		},
	}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/solicitudes-tutoria",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetSolicitudesTutoria(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200 y obtuvo %d", rec.Code)
	}
}