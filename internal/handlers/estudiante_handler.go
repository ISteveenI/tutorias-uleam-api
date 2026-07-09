package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

var Estudiantes = storage.NewEstudianteStorage()

// Crear estudiante
func CreateEstudiante(w http.ResponseWriter, r *http.Request) {

	var estudiante models.Estudiante

	err := json.NewDecoder(r.Body).Decode(&estudiante)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if estudiante.Nombres == "" {
		http.Error(w, "Los nombres son obligatorios", http.StatusBadRequest)
		return
	}

	if estudiante.Apellidos == "" {
		http.Error(w, "Los apellidos son obligatorios", http.StatusBadRequest)
		return
	}

	if estudiante.Correo == "" {
		http.Error(w, "El correo es obligatorio", http.StatusBadRequest)
		return
	}

	if estudiante.Carrera == "" {
		http.Error(w, "La carrera es obligatoria", http.StatusBadRequest)
		return
	}

	nuevo := Estudiantes.Create(estudiante)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nuevo)
}

// Obtener todos los estudiantes
func GetEstudiantes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Estudiantes.GetAll())
}

// Obtener estudiante por ID
func GetEstudianteByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	estudiante, err := Estudiantes.GetByID(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estudiante)
}

// Actualizar estudiante
func UpdateEstudiante(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var estudiante models.Estudiante

	err = json.NewDecoder(r.Body).Decode(&estudiante)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if estudiante.Nombres == "" {
		http.Error(w, "Los nombres son obligatorios", http.StatusBadRequest)
		return
	}

	if estudiante.Apellidos == "" {
		http.Error(w, "Los apellidos son obligatorios", http.StatusBadRequest)
		return
	}

	actualizado, err := Estudiantes.Update(uint(id), estudiante)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// Eliminar estudiante
func DeleteEstudiante(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = Estudiantes.Delete(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Estudiante eliminado correctamente",
	})
}