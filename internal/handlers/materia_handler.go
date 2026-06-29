package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)


// Crear una nueva materia.
func CreateMateria(w http.ResponseWriter, r *http.Request) {

	var materia models.Materia
	err := json.NewDecoder(r.Body).Decode(&materia)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// Validaciones
	if materia.Nombre == "" {
		http.Error(w, "El nombre de la materia es obligatorio", http.StatusBadRequest)
		return
	}
	if materia.Codigo == "" {
		http.Error(w, "El código de la materia es obligatorio", http.StatusBadRequest)
		return
	}
	materia = storage.MateriaRepo.Create(materia)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(materia)

}

// Obtener todas las materias.
func GetMaterias(w http.ResponseWriter, r *http.Request) {
	materias := storage.MateriaRepo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(materias)

}

// Obtener materia por ID.
func GetMateriaByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	materia, err := storage.MateriaRepo.GetByID(id)

	if err != nil {
		http.Error(
			w,
			"Materia no encontrada",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(materia)

}

// Actualizar una materia.
func UpdateMateria(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	var materia models.Materia
	err = json.NewDecoder(r.Body).Decode(&materia)
	if err != nil {

		http.Error(w, "Datos inválidos", http.StatusBadRequest)

		return
	}

	// Validaciones

	if materia.Nombre == "" {

		http.Error(
			w,
			"El nombre de la materia es obligatorio",
			http.StatusBadRequest,
		)

		return
	}

	updated, err := storage.MateriaRepo.Update(id, materia)

	if err != nil {

		http.Error(
			w,
			"Materia no encontrada",
			http.StatusNotFound,
		)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)

}

// Eliminar una materia.
func DeleteMateria(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	err = storage.MateriaRepo.Delete(id)

	if err != nil {
		http.Error(
			w,
			"Materia no encontrada",
			http.StatusNotFound,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Materia eliminada correctamente",

	})

}