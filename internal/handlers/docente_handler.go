package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

// Crear docente
func CreateDocente(w http.ResponseWriter, r *http.Request) {

	var docente models.Docente

	err := json.NewDecoder(r.Body).Decode(&docente)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// VALIDACIONES
	if docente.Nombres == "" {
		http.Error(w, "Nombres obligatorios", http.StatusBadRequest)
		return
	}

	if docente.Apellidos == "" {
		http.Error(w, "Apellidos obligatorios", http.StatusBadRequest)
		return
	}

	if docente.Correo == "" {
		http.Error(w, "Correo obligatorio", http.StatusBadRequest)
		return
	}

	if docente.Departamento == "" {
		http.Error(w, "Departamento obligatorio", http.StatusBadRequest)
		return
	}

	if docente.Especialidad == "" {
		http.Error(w, "Especialidad obligatoria", http.StatusBadRequest)
		return
	}

	// CREAR
	docente = storage.DocenteRepo.Create(docente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(docente)
}

// Obtener todos los docentes
func GetDocentes(w http.ResponseWriter, r *http.Request) {

	docentes := storage.DocenteRepo.GetAll()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(docentes)
}

// Obtener docente por ID
func GetDocenteByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	docente, err := storage.DocenteRepo.GetByID(id)

	if err != nil {
		http.Error(w, "Docente no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(docente)
}

// Actualizar docente
func UpdateDocente(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var docente models.Docente

	err = json.NewDecoder(r.Body).Decode(&docente)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// VALIDACIONES
	if docente.Nombres == "" {
		http.Error(w, "Nombres obligatorios", http.StatusBadRequest)
		return
	}

	if docente.Correo == "" {
		http.Error(w, "Correo obligatorio", http.StatusBadRequest)
		return
	}

	updated, err := storage.DocenteRepo.Update(id, docente)

	if err != nil {
		http.Error(w, "Docente no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(updated)
}

// Eliminar docente
func DeleteDocente(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = storage.DocenteRepo.Delete(id)

	if err != nil {
		http.Error(w, "Docente no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Docente eliminado correctamente",
	})
}