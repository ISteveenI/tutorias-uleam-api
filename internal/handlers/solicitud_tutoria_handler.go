package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
)

type SolicitudTutoriaHandler struct {
	service services.SolicitudTutoriaServiceInterface
}

func NewSolicitudTutoriaHandler(service services.SolicitudTutoriaServiceInterface) *SolicitudTutoriaHandler {
	return &SolicitudTutoriaHandler{
		service: service,
	}
}

func (h *SolicitudTutoriaHandler) CreateSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	var solicitud models.SolicitudTutoria

	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(context.Background(), &solicitud); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(solicitud)
}

func (h *SolicitudTutoriaHandler) GetSolicitudesTutoria(w http.ResponseWriter, r *http.Request) {

	solicitudes, err := h.service.GetAll(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitudes)
}

func (h *SolicitudTutoriaHandler) GetSolicitudByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	solicitud, err := h.service.GetByID(context.Background(), uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitud)
}

func (h *SolicitudTutoriaHandler) UpdateSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var solicitud models.SolicitudTutoria

	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(context.Background(), uint(id), &solicitud); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitud)
}

func (h *SolicitudTutoriaHandler) DeleteSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(context.Background(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Solicitud de tutoría eliminada correctamente",
	})
}