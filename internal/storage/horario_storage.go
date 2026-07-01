package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type HorarioStorage struct {
	mu       sync.Mutex
	horarios map[int]models.HorarioDocente
	nextID   int
}

func NewHorarioStorage() *HorarioStorage {
	return &HorarioStorage{
		horarios: make(map[int]models.HorarioDocente),
		nextID:   1,
	}
}

// Validar cruce de horario
func (h *HorarioStorage) ExisteCruceHorario(nuevo models.HorarioDocente) bool {

	for _, horario := range h.horarios {
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

func (h *HorarioStorage) Create(horario models.HorarioDocente) (models.HorarioDocente, error) {

	h.mu.Lock()
	defer h.mu.Unlock()
	if h.ExisteCruceHorario(horario) {
		return models.HorarioDocente{},
			errors.New("existe cruce de horario")
	}
	horario.ID = h.nextID
	h.horarios[horario.ID] = horario
	h.nextID++
	return horario, nil
}

func (h *HorarioStorage) GetAll() []models.HorarioDocente {

	h.mu.Lock()
	defer h.mu.Unlock()
	lista := make([]models.HorarioDocente, 0, len(h.horarios))
	for _, horario := range h.horarios {
		lista = append(lista, horario)
	}
	return lista
}

func (h *HorarioStorage) GetByID(id int) (models.HorarioDocente, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	horario, exists := h.horarios[id]
	if !exists {
		return models.HorarioDocente{},
			errors.New("horario no encontrado")
	}
	return horario, nil
}

func (h *HorarioStorage) Update(id int, horario models.HorarioDocente) (models.HorarioDocente, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	_, exists := h.horarios[id]
	if !exists {
		return models.HorarioDocente{},
			errors.New("horario no encontrado")
	}
	horario.ID = id
	h.horarios[id] = horario
	return horario, nil
}

func (h *HorarioStorage) Delete(id int) error {

	h.mu.Lock()
	defer h.mu.Unlock()
	_, exists := h.horarios[id]
	if !exists {
		return errors.New("horario no encontrado")
	}
	delete(h.horarios, id)
	return nil
}

var HorarioRepo = NewHorarioStorage()
