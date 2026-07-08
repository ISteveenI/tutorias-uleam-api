package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type fakeSesionService struct {
	sesiones map[uint]models.SesionTutoria
	nextID   uint
}

func newFakeSesionService() *fakeSesionService {
	return &fakeSesionService{
		sesiones: make(map[uint]models.SesionTutoria),
		nextID:   1,
	}
}

func (f *fakeSesionService) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	sesion.ID = f.nextID
	f.sesiones[sesion.ID] = *sesion
	f.nextID++
	return nil
}

func (f *fakeSesionService) GetByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	sesion := f.sesiones[id]
	return &sesion, nil
}

func (f *fakeSesionService) GetAll(ctx context.Context) ([]models.SesionTutoria, error) {
	sesiones := make([]models.SesionTutoria, 0, len(f.sesiones))

	for _, sesion := range f.sesiones {
		sesiones = append(sesiones, sesion)
	}

	return sesiones, nil
}

func (f *fakeSesionService) Update(ctx context.Context, id uint, sesion *models.SesionTutoria) error {
	sesion.ID = id
	f.sesiones[id] = *sesion
	return nil
}

func (f *fakeSesionService) Delete(ctx context.Context, id uint) error {
	delete(f.sesiones, id)
	return nil
}

func (f *fakeSesionService) RegistrarAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	return nil
}

func (f *fakeSesionService) SubirEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	return nil
}

func TestCreateSesionSinTokenResponde401(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	sesionHandler.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusTeapot {
		t.Fatalf("se esperaba status 401, pero se obtuvo %d", rr.Code)
	}
}

func TestCreateSesionConTokenResponde201YGuardaEnFake(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"solicitud_id":  1,
		"fecha_sesion":  "2026-07-01",
		"hora_inicio":   "09:00",
		"hora_fin":      "10:00",
		"observaciones": "Primera sesion de tutoria",
		"estado":        "Programada",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := httptest.NewRecorder()

	sesionHandler.Routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, pero se obtuvo %d. Body: %s", rr.Code, rr.Body.String())
	}

	if len(fakeService.sesiones) != 1 {
		t.Fatalf("se esperaba que el fake guarde 1 sesion, pero guardo %d", len(fakeService.sesiones))
	}
}
