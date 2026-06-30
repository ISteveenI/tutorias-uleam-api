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

type SesionHandler struct {
	service services.SesionServiceInterface
}

func NewSesionHandler(service services.SesionServiceInterface) *SesionHandler {
	return &SesionHandler{
		service: service,
	}
}

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

func (h *SesionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sesiones, err := h.service.GetAll(r.Context())
	if err != nil {
		escribirSesionError(w, http.StatusInternalServerError, "No se pudieron obtener las sesiones")
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesiones)
}

func (h *SesionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
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

func (h *SesionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
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

func (h *SesionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
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

func (h *SesionHandler) CreateAsistencia(w http.ResponseWriter, r *http.Request) {
	sesionID, err := obtenerSesionID(r)
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

func (h *SesionHandler) CreateEvidencia(w http.ResponseWriter, r *http.Request) {
	sesionID, err := obtenerSesionID(r)
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

func obtenerSesionID(r *http.Request) (uint, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	return uint(id), err
}

func escribirSesionJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func escribirSesionError(w http.ResponseWriter, status int, mensaje string) {
	escribirSesionJSON(w, status, map[string]string{
		"error": mensaje,
	})
}

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
