package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type SesionStorage struct {
	mu       sync.Mutex
	sesiones map[int]models.SesionTutoria
	nextID   int
}

func NewSesionStorage() *SesionStorage {
	return &SesionStorage{
		sesiones: make(map[int]models.SesionTutoria),
		nextID:   1,
	}
}

func (s *SesionStorage) Create(sesion models.SesionTutoria) models.SesionTutoria {
	s.mu.Lock()
	defer s.mu.Unlock()

	sesion.ID = s.nextID
	s.sesiones[sesion.ID] = sesion
	s.nextID++

	return sesion
}

func (s *SesionStorage) GetAll() []models.SesionTutoria {
	s.mu.Lock()
	defer s.mu.Unlock()

	lista := make([]models.SesionTutoria, 0, len(s.sesiones))
	for _, sesion := range s.sesiones {
		lista = append(lista, sesion)
	}
	return lista
}

func (s *SesionStorage) GetByID(id int) (models.SesionTutoria, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sesion, exists := s.sesiones[id]
	if !exists {
		return models.SesionTutoria{}, errors.New("sesion no encontrada")
	}

	return sesion, nil
}

func (s *SesionStorage) Update(id int, sesion models.SesionTutoria) (models.SesionTutoria, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.sesiones[id]
	if !exists {
		return models.SesionTutoria{}, errors.New("sesion no encontrada")
	}

	sesion.ID = id
	s.sesiones[id] = sesion
	return sesion, nil
}

func (s *SesionStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.sesiones[id]
	if !exists {
		return errors.New("sesion no encontrada")
	}

	delete(s.sesiones, id)
	return nil
}
