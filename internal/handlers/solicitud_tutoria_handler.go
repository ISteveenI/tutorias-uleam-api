package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

func CreateSolicitudTutoria(w http.ResponseWriter, r *http.Request) {

	var solicitud models.SolicitudTutoria

	err := json.NewDecoder(r.Body).Decode(&solicitud)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	solicitud.ID = len(storage.SolicitudesTutoria) + 1

	storage.SolicitudesTutoria = append(
		storage.SolicitudesTutoria,
		solicitud,
	)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(solicitud)
}

func SolicitudesTutoria(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(storage.SolicitudesTutoria)
}
