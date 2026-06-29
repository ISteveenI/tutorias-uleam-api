package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)


// Verifica si existe cruce de horario.
func ExisteCruceHorario(
	nuevo models.HorarioDocente,
) bool {

	horarios := storage.HorarioRepo.GetAll()

	for _, horario := range horarios {

		if horario.DocenteID != nuevo.DocenteID {
			continue
		}

		if horario.DiaSemana != nuevo.DiaSemana {
			continue
		}

		if nuevo.HoraInicio < horario.HoraFin &&
			nuevo.HoraFin > horario.HoraInicio {

			return true
		}
	}

	return false
}

// Crear nuevo horario docente.
func CreateHorario(w http.ResponseWriter, r *http.Request) {

	var horario models.HorarioDocente

	err := json.NewDecoder(r.Body).Decode(&horario)

	if err != nil {

		http.Error(w, "Datos inválidos", http.StatusBadRequest)

		return
	}

	// Validaciones

	if horario.DocenteID <= 0 {

		http.Error(
			w,
			"El ID del docente es obligatorio",
			http.StatusBadRequest,
		)

		return
	}

	if horario.MateriaID <= 0 {

		http.Error(
			w,
			"El ID de la materia es obligatorio",
			http.StatusBadRequest,
		)

		return
	}

	if horario.DiaSemana == "" {

		http.Error(
			w,
			"El día de la semana es obligatorio",
			http.StatusBadRequest,
		)

		return
	}

	if horario.HoraInicio == "" {

		http.Error(
			w,
			"La hora de inicio es obligatoria",
			http.StatusBadRequest,
		)

		return
	}

	if horario.HoraFin == "" {
		http.Error(
			w,
			"La hora de fin es obligatoria",
			http.StatusBadRequest,
		)
		return
	}

	// Validar cruce horario

	if ExisteCruceHorario(horario) {

		http.Error(
			w,
			"Existe un cruce de horario para este docente",
			http.StatusBadRequest,
		)

		return
	}

	horario, err = storage.HorarioRepo.Create(horario)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(horario)

}

// Obtener todos los horarios.
func GetHorarios(w http.ResponseWriter, r *http.Request) {

	horarios := storage.HorarioRepo.GetAll()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(horarios)

}

// Obtener horario por ID.
func GetHorarioByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	horario, err := storage.HorarioRepo.GetByID(id)

	if err != nil {

		http.Error(
			w,
			"Horario no encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(horario)

}

// Actualizar horario.
func UpdateHorario(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	var horario models.HorarioDocente

	err = json.NewDecoder(r.Body).Decode(&horario)

	if err != nil {

		http.Error(w, "Datos inválidos", http.StatusBadRequest)

		return
	}

	updated, err := storage.HorarioRepo.Update(id, horario)

	if err != nil {
		http.Error(
			w,
			"Horario no encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)

}

// Eliminar horario.
func DeleteHorario(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	err = storage.HorarioRepo.Delete(id)
	if err != nil {

		http.Error(
			w,
			"Horario no encontrado",
			http.StatusNotFound,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Horario eliminado correctamente",

	})

}