package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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

func GetSolicitudesTutoria(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(storage.SolicitudesTutoria)
}

func GetSolicitudByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, solicitud := range storage.SolicitudesTutoria {

		if solicitud.ID == id {

			json.NewEncoder(w).Encode(solicitud)
			return
		}
	}

	http.Error(w, "Solicitud no encontrada", http.StatusNotFound)
}
