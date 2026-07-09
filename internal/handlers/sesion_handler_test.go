package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeSesionService struct {
	sesiones    map[uint]models.SesionTutoria
	asistencias []models.Asistencia
	evidencias  []models.Evidencia
	nextID      uint

	errCreate              error
	errGetByID             error
	errGetAll              error
	errUpdate              error
	errDelete              error
	errRegistrarAsistencia error
	errSubirEvidencia      error
}

func newFakeSesionService() *fakeSesionService {
	return &fakeSesionService{
		sesiones: make(map[uint]models.SesionTutoria),
		nextID:   1,
	}
}

func (f *fakeSesionService) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	if f.errCreate != nil {
		return f.errCreate
	}

	sesion.ID = f.nextID
	f.sesiones[sesion.ID] = *sesion
	f.nextID++

	return nil
}

func (f *fakeSesionService) GetByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	if f.errGetByID != nil {
		return nil, f.errGetByID
	}

	sesion, existe := f.sesiones[id]
	if !existe {
		return nil, repositories.ErrRegistroNoEncontrado
	}

	return &sesion, nil
}

func (f *fakeSesionService) GetAll(ctx context.Context) ([]models.SesionTutoria, error) {
	if f.errGetAll != nil {
		return nil, f.errGetAll
	}

	sesiones := make([]models.SesionTutoria, 0, len(f.sesiones))
	for _, sesion := range f.sesiones {
		sesiones = append(sesiones, sesion)
	}

	return sesiones, nil
}

func (f *fakeSesionService) Update(ctx context.Context, id uint, sesion *models.SesionTutoria) error {
	if f.errUpdate != nil {
		return f.errUpdate
	}

	sesion.ID = id
	f.sesiones[id] = *sesion

	return nil
}

func (f *fakeSesionService) Delete(ctx context.Context, id uint) error {
	if f.errDelete != nil {
		return f.errDelete
	}

	if _, existe := f.sesiones[id]; !existe {
		return repositories.ErrRegistroNoEncontrado
	}

	delete(f.sesiones, id)
	return nil
}

func (f *fakeSesionService) RegistrarAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	if f.errRegistrarAsistencia != nil {
		return f.errRegistrarAsistencia
	}

	f.asistencias = append(f.asistencias, *asistencia)
	return nil
}

func (f *fakeSesionService) SubirEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	if f.errSubirEvidencia != nil {
		return f.errSubirEvidencia
	}

	f.evidencias = append(f.evidencias, *evidencia)
	return nil
}

func nuevaSesionParaHandlerTest(id uint) models.SesionTutoria {
	return models.SesionTutoria{
		ID:            id,
		SolicitudID:   1,
		FechaSesion:   "2026-07-01",
		HoraInicio:    "09:00",
		HoraFin:       "10:00",
		Observaciones: "Sesion de prueba",
		Estado:        "Programada",
	}
}

func nuevaPeticionJSON(t *testing.T, metodo string, ruta string, body any, conToken bool) *http.Request {
	t.Helper()

	var lector *bytes.Reader

	if body == nil {
		lector = bytes.NewReader(nil)
	} else {
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)
		lector = bytes.NewReader(jsonBody)
	}

	req := httptest.NewRequest(metodo, ruta, lector)
	req.Header.Set("Content-Type", "application/json")

	if conToken {
		req.Header.Set("Authorization", "Bearer token-prueba")
	}

	return req
}

func ejecutarRequest(router http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestCreateSesionSinTokenResponde401(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	sesionHandler.Routes().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
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

	req := nuevaPeticionJSON(t, http.MethodPost, "/", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Len(t, fakeService.sesiones, 1)
}

func TestCreateSesionJSONInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{"))
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateSesionDatosInvalidosResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.errCreate = services.ErrDatosInvalidos

	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"solicitud_id": 1,
		"fecha_sesion": "2026-07-01",
		"hora_inicio":  "09:00",
		"hora_fin":     "10:00",
	}

	req := nuevaPeticionJSON(t, http.MethodPost, "/", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllSesionesResponde200(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.sesiones[1] = nuevaSesionParaHandlerTest(1)

	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	require.Equal(t, http.StatusOK, rr.Code)

	var respuesta []models.SesionTutoria
	err := json.NewDecoder(rr.Body).Decode(&respuesta)
	require.NoError(t, err)

	assert.Len(t, respuesta, 1)
	assert.Equal(t, uint(1), respuesta[0].ID)
}

func TestGetAllServicioErrorResponde500(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.errGetAll = errors.New("fallo al listar sesiones")

	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetByIDSesionResponde200(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.sesiones[1] = nuevaSesionParaHandlerTest(1)

	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	require.Equal(t, http.StatusOK, rr.Code)

	var respuesta models.SesionTutoria
	err := json.NewDecoder(rr.Body).Decode(&respuesta)
	require.NoError(t, err)

	assert.Equal(t, uint(1), respuesta.ID)
	assert.Equal(t, "Programada", respuesta.Estado)
}

func TestGetByIDInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetByIDNoEncontradoResponde404(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodGet, "/999", nil)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUpdateSesionConTokenResponde200YActualiza(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.sesiones[1] = nuevaSesionParaHandlerTest(1)

	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"solicitud_id":  1,
		"fecha_sesion":  "2026-07-02",
		"hora_inicio":   "11:00",
		"hora_fin":      "12:00",
		"observaciones": "Sesion actualizada",
		"estado":        "Reprogramada",
	}

	req := nuevaPeticionJSON(t, http.MethodPut, "/1", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	require.Equal(t, http.StatusOK, rr.Code)

	actualizada := fakeService.sesiones[1]
	assert.Equal(t, uint(1), actualizada.ID)
	assert.Equal(t, "Reprogramada", actualizada.Estado)
	assert.Equal(t, "Sesion actualizada", actualizada.Observaciones)
}

func TestUpdateSesionSinTokenResponde401(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"solicitud_id": 1,
		"fecha_sesion": "2026-07-02",
		"hora_inicio":  "11:00",
		"hora_fin":     "12:00",
	}

	req := nuevaPeticionJSON(t, http.MethodPut, "/1", body, false)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestUpdateSesionJSONInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBufferString("{"))
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateServicioErrorInternoResponde500(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.errUpdate = errors.New("fallo interno")

	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"solicitud_id": 1,
		"fecha_sesion": "2026-07-02",
		"hora_inicio":  "11:00",
		"hora_fin":     "12:00",
	}

	req := nuevaPeticionJSON(t, http.MethodPut, "/1", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteSesionConTokenResponde200YElimina(t *testing.T) {
	fakeService := newFakeSesionService()
	fakeService.sesiones[1] = nuevaSesionParaHandlerTest(1)

	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodDelete, "/1", nil)
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Empty(t, fakeService.sesiones)
}

func TestDeleteSesionNoEncontradaResponde404(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodDelete, "/999", nil)
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestCreateAsistenciaConTokenResponde201YAsignaSesionID(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"estudiante_asistio": true,
		"docente_asistio":    true,
		"observacion":        "Asistieron ambos",
	}

	req := nuevaPeticionJSON(t, http.MethodPost, "/5/asistencias", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	require.Equal(t, http.StatusCreated, rr.Code)
	require.Len(t, fakeService.asistencias, 1)

	assert.Equal(t, uint(5), fakeService.asistencias[0].SesionID)
	assert.True(t, fakeService.asistencias[0].EstudianteAsistio)
	assert.True(t, fakeService.asistencias[0].DocenteAsistio)
}

func TestCreateAsistenciaIDInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"estudiante_asistio": true,
		"docente_asistio":    true,
	}

	req := nuevaPeticionJSON(t, http.MethodPost, "/abc/asistencias", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateAsistenciaJSONInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPost, "/1/asistencias", bytes.NewBufferString("{"))
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateEvidenciaConTokenResponde201YAsignaSesionID(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	body := map[string]any{
		"tipo_archivo": "pdf",
		"archivo_url":  "https://example.com/evidencia.pdf",
		"descripcion":  "Evidencia de la tutoria",
	}

	req := nuevaPeticionJSON(t, http.MethodPost, "/7/evidencias", body, true)
	rr := ejecutarRequest(sesionHandler.Routes(), req)

	require.Equal(t, http.StatusCreated, rr.Code)
	require.Len(t, fakeService.evidencias, 1)

	assert.Equal(t, uint(7), fakeService.evidencias[0].SesionID)
	assert.Equal(t, "pdf", fakeService.evidencias[0].TipoArchivo)
	assert.Equal(t, "https://example.com/evidencia.pdf", fakeService.evidencias[0].ArchivoURL)
}

func TestCreateEvidenciaJSONInvalidoResponde400(t *testing.T) {
	fakeService := newFakeSesionService()
	sesionHandler := handlers.NewSesionHandler(fakeService)

	req := httptest.NewRequest(http.MethodPost, "/1/evidencias", bytes.NewBufferString("{"))
	req.Header.Set("Authorization", "Bearer token-prueba")

	rr := ejecutarRequest(sesionHandler.Routes(), req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
