package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
)

// SesionHandler recibe las peticiones HTTP relacionadas con sesiones de tutoría.
// No contiene lógica de negocio ni acceso directo a base de datos, esas tareas
// se delegan al service.
type SesionHandler struct {
	service services.SesionServiceInterface
}

// NewSesionHandler crea un handler de sesiones usando la interfaz del service.
// Esto permite probar el handler con un fake en memoria durante los tests.
func NewSesionHandler(service services.SesionServiceInterface) *SesionHandler {
	return &SesionHandler{
		service: service,
	}
}

// Routes define las rutas HTTP del módulo de sesiones.
// Las rutas de consulta son públicas, mientras que crear, actualizar,
// eliminar, registrar asistencia y subir evidencia requieren autenticación.
func (h *SesionHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)

	r.Group(func(protegida chi.Router) {
		protegida.Use(middlewares.RequireAuth)

		protegida.Post("/", h.Create)
		protegida.Put("/{id}", h.Update)
		protegida.Delete("/{id}", h.Delete)

		protegida.Post("/{id}/asistencias", h.CreateAsistencia)
		protegida.Post("/{id}/evidencias", h.CreateEvidencia)
	})

	return r
}

// Create recibe el JSON de una sesión y solicita al service su creación.
// Si el JSON es inválido o no cumple las reglas del service, responde con error.
func (h *SesionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var sesion models.SesionTutoria

	if err := json.NewDecoder(r.Body).Decode(&sesion); err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if err := h.service.Create(r.Context(), &sesion); err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusCreated, sesion)
}

// GetAll lista todas las sesiones de tutoría registradas.
func (h *SesionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sesiones, err := h.service.GetAll(r.Context())
	if err != nil {
		escribirSesionError(w, http.StatusInternalServerError, "No se pudieron obtener las sesiones")
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesiones)
}

// GetByID obtiene una sesión específica usando el id recibido en la URL.
func (h *SesionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	sesion, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesion)
}

// Update actualiza una sesión existente. El id se obtiene desde la URL
// y los nuevos datos se reciben desde el cuerpo JSON.
func (h *SesionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var sesion models.SesionTutoria
	if err := json.NewDecoder(r.Body).Decode(&sesion); err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if err := h.service.Update(r.Context(), id, &sesion); err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesion)
}

// Delete elimina una sesión de tutoría usando el id recibido en la URL.
func (h *SesionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusOK, map[string]string{
		"mensaje": "Sesion eliminada correctamente",
	})
}

// CreateAsistencia registra la asistencia de una sesión.
// El id de la sesión viene desde la URL y se asigna al modelo antes de enviarlo al service.
func (h *SesionHandler) CreateAsistencia(w http.ResponseWriter, r *http.Request) {
	sesionID, err := idDeURL(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var asistencia models.Asistencia
	if err := json.NewDecoder(r.Body).Decode(&asistencia); err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	asistencia.SesionID = sesionID

	if err := h.service.RegistrarAsistencia(r.Context(), &asistencia); err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusCreated, asistencia)
}

// CreateEvidencia registra una evidencia asociada a una sesión.
// El id de la sesión se toma desde la URL y no desde el JSON del cliente.
func (h *SesionHandler) CreateEvidencia(w http.ResponseWriter, r *http.Request) {
	sesionID, err := idDeURL(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var evidencia models.Evidencia
	if err := json.NewDecoder(r.Body).Decode(&evidencia); err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	evidencia.SesionID = sesionID

	if err := h.service.SubirEvidencia(r.Context(), &evidencia); err != nil {
		manejarErrorSesion(w, err)
		return
	}

	escribirSesionJSON(w, http.StatusCreated, evidencia)
}

// idDeURL centraliza la lectura y conversión del parámetro "id" desde la URL.
// Este refactor aplica DRY: evita repetir chi.URLParam y strconv.ParseUint
// en cada método del handler.
func idDeURL(r *http.Request) (uint, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	return uint(id), err
}

// escribirSesionJSON centraliza la escritura de respuestas JSON.
// Así todos los endpoints responden con el mismo Content-Type y estructura.
func escribirSesionJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// escribirSesionError estandariza las respuestas de error del módulo.
func escribirSesionError(w http.ResponseWriter, status int, mensaje string) {
	escribirSesionJSON(w, status, map[string]string{
		"error": mensaje,
	})
}

// manejarErrorSesion traduce errores del service o repository a códigos HTTP.
// Esto evita repetir el mismo switch de errores en cada endpoint.
func manejarErrorSesion(w http.ResponseWriter, err error) {
	if errors.Is(err, services.ErrDatosInvalidos) {
		escribirSesionError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}

	if errors.Is(err, repositories.ErrRegistroNoEncontrado) {
		escribirSesionError(w, http.StatusNotFound, "Sesion no encontrada")
		return
	}

	escribirSesionError(w, http.StatusInternalServerError, "Error interno del servidor")
}
