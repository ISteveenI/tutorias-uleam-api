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

	// Validaciones
	if disponibilidad.Materia == "" {
		http.Error(w, "La materia es obligatoria", http.StatusBadRequest)
		return
	}

	if disponibilidad.Cupos <= 0 {
		http.Error(w, "Los cupos deben ser mayores a cero", http.StatusBadRequest)
		return
	}

	if disponibilidad.DocenteID <= 0 {
		http.Error(w, "El ID del docente es obligatorio", http.StatusBadRequest)
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

func UpdateDisponibilidad(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var updated models.DisponibilidadDocente
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	for i, disponibilidad := range storage.Disponibilidades {
		if disponibilidad.ID == id {
			updated.ID = id
			storage.Disponibilidades[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Disponibilidad no encontrada", http.StatusNotFound)
}

func DeleteDisponibilidad(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for i, disponibilidad := range storage.Disponibilidades {
		if disponibilidad.ID == id {
			storage.Disponibilidades = append(
				storage.Disponibilidades[:i],
				storage.Disponibilidades[i+1:]...,
			)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Disponibilidad eliminada"))
			return
		}
	}
	http.Error(w, "Disponibilidad no encontrada", http.StatusNotFound)
}