package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

var SolicitudesTutoria = storage.NewSolicitudTutoriaStorage()

// Crea una nueva solicitud de tutoría.
func CreateSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	var solicitud models.SolicitudTutoria

	err := json.NewDecoder(r.Body).Decode(&solicitud)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validaciones
	if solicitud.EstudianteID <= 0 {
		http.Error(w, "El estudiante es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.HorarioDocenteID <= 0 {
		http.Error(w, "El horario docente es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.TipoTutoriaID <= 0 {
		http.Error(w, "El tipo de tutoría es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.Tema == "" {
		http.Error(w, "El tema es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.FechaSolicitud == "" {
		http.Error(w, "La fecha de la solicitud es obligatoria", http.StatusBadRequest)
		return
	}

	if solicitud.Estado == "" {
		solicitud.Estado = "Pendiente"
	}

	solicitud = SolicitudesTutoria.Create(solicitud)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(solicitud)
}

// Obtiene todas las solicitudes de tutoría.
func GetSolicitudesTutoria(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(SolicitudesTutoria.GetAll())
}

// Obtiene una solicitud por ID.
func GetSolicitudByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	solicitud, err := SolicitudesTutoria.GetByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(solicitud)
}

// Actualiza una solicitud de tutoría.
func UpdateSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var solicitud models.SolicitudTutoria

	err = json.NewDecoder(r.Body).Decode(&solicitud)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if solicitud.EstudianteID <= 0 {
		http.Error(w, "El estudiante es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.HorarioDocenteID <= 0 {
		http.Error(w, "El horario docente es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.TipoTutoriaID <= 0 {
		http.Error(w, "El tipo de tutoría es obligatorio", http.StatusBadRequest)
		return
	}

	if solicitud.Tema == "" {
		http.Error(w, "El tema es obligatorio", http.StatusBadRequest)
		return
	}

	actualizada, err := SolicitudesTutoria.Update(id, solicitud)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(actualizada)
}

// Elimina una solicitud de tutoría.
func DeleteSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = SolicitudesTutoria.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Solicitud de tutoría eliminada correctamente",
	})
}