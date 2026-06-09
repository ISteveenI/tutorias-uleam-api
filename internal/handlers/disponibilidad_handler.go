package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func GetDisponibilidadByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)	
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, disponibilidad := range storage.Disponibilidades {
		if disponibilidad.ID == id {
			json.NewEncoder(w).Encode(disponibilidad)
			return
		}
	}	
	http.Error(w, "Disponibilidad no encontrada", http.StatusNotFound)
}