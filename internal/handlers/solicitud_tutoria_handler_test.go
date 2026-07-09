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

	if s.Tema == "" {
		return errors.New("datos inválidos")
	}

	s.ID = 1
	f.data = append(f.data, *s)

	return nil
}

func (f *fakeSolicitudService) GetByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {

	for _, v := range f.data {

		if uint(v.ID) == id {
			return &v, nil
		}
	}

	return nil, errors.New("solicitud no encontrada")
}

func (f *fakeSolicitudService) GetAll(ctx context.Context) ([]models.SolicitudTutoria, error) {
	return f.data, nil
}

func (f *fakeSolicitudService) Update(ctx context.Context, id uint, s *models.SolicitudTutoria) error {

	if id == 999 {
		return errors.New("no encontrada")
	}

	if s.Tema == "" {
		return errors.New("datos inválidos")
	}

	return nil
}

func (f *fakeSolicitudService) Delete(ctx context.Context, id uint) error {

	if id == 999 {
		return errors.New("no encontrada")
	}

	return nil
}

func TestCreateSolicitudHandler_OK(t *testing.T) {

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
}

func TestCreateSolicitudHandler_JSONInvalido(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/solicitudes-tutoria",
		bytes.NewBufferString("{"),
	)

	rec := httptest.NewRecorder()

	handler.CreateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestCreateSolicitudHandler_ErrorServicio(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	body := models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "",
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

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestGetSolicitudesHandler_OK(t *testing.T) {

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

func TestGetSolicitudByID_OK(t *testing.T) {

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
		"/solicitudes/1",
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.GetSolicitudByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200 y obtuvo %d", rec.Code)
	}
}

func TestGetSolicitudByID_NoExiste(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodGet,
		"/solicitudes/999",
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.GetSolicitudByID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 y obtuvo %d", rec.Code)
	}
}

func TestGetSolicitudByID_IDInvalido(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodGet,
		"/solicitudes/abc",
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.GetSolicitudByID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestUpdateSolicitud_OK(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	body := models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 1,
		TipoTutoriaID:    1,
		Tema:             "Nuevo Tema",
		FechaSolicitud:   "2026-06-30",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPut,
		"/solicitudes/1",
		bytes.NewReader(jsonBody),
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.UpdateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200 y obtuvo %d", rec.Code)
	}
}

func TestUpdateSolicitud_JSONInvalido(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodPut,
		"/solicitudes/1",
		bytes.NewBufferString("{"),
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.UpdateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestUpdateSolicitud_IDInvalido(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodPut,
		"/solicitudes/abc",
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.UpdateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestUpdateSolicitud_NoExiste(t *testing.T) {

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
		http.MethodPut,
		"/solicitudes/999",
		bytes.NewReader(jsonBody),
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.UpdateSolicitudTutoria(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}

func TestDeleteSolicitud_NoExiste(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/solicitudes/999",
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")

	req = req.WithContext(
		context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
		),
	)

	rec := httptest.NewRecorder()

	handler.DeleteSolicitudTutoria(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 y obtuvo %d", rec.Code)
	}
}

func TestRoutes_Create(t *testing.T) {

	fake := &fakeSolicitudService{}
	handler := NewSolicitudTutoriaHandler(fake)

	router := handler.Routes()

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
		"/",
		bytes.NewReader(jsonBody),
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201 y obtuvo %d", rec.Code)
	}
}

func TestRoutes_GetAll(t *testing.T) {

	fake := &fakeSolicitudService{}
	handler := NewSolicitudTutoriaHandler(fake)

	router := handler.Routes()

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200 y obtuvo %d", rec.Code)
	}
}

func TestRoutes_NotFound(t *testing.T) {

	fake := &fakeSolicitudService{}

	handler := NewSolicitudTutoriaHandler(fake)

	router := handler.Routes()

	req := httptest.NewRequest(
		http.MethodGet,
		"/ruta-inexistente",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// Chi interpreta "/ruta-inexistente" como {id},
	// luego el handler intenta convertirlo a uint
	// y responde 400 (BadRequest).

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 y obtuvo %d", rec.Code)
	}
}
