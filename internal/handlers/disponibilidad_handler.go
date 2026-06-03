package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

func CreateDisponibilidad(w http.ResponseWriter, r *http.Request) {

	var disponibilidad models.DisponibilidadDocente

	err := json.NewDecoder(r.Body).Decode(&disponibilidad)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	disponibilidad.ID = len(storage.Disponibilidades) + 1

	storage.Disponibilidades = append(
		storage.Disponibilidades,
		disponibilidad,
	)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(disponibilidad)
}

func GetDisponibilidades(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(storage.Disponibilidades)
}