package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type SolicitudTutoriaStorage struct {
	mu           sync.Mutex
	solicitudes  map[int]models.SolicitudTutoria
	nextID       int
}

func NewSolicitudTutoriaStorage() *SolicitudTutoriaStorage {
	return &SolicitudTutoriaStorage{
		solicitudes: make(map[int]models.SolicitudTutoria),
		nextID:      1,
	}
}

func (s *SolicitudTutoriaStorage) Create(
	solicitud models.SolicitudTutoria,
) models.SolicitudTutoria {

	s.mu.Lock()
	defer s.mu.Unlock()

	solicitud.ID = s.nextID
	s.solicitudes[solicitud.ID] = solicitud
	s.nextID++

	return solicitud
}

func (s *SolicitudTutoriaStorage) GetAll() []models.SolicitudTutoria {

	s.mu.Lock()
	defer s.mu.Unlock()

	lista := make([]models.SolicitudTutoria, 0, len(s.solicitudes))

	for _, solicitud := range s.solicitudes {
		lista = append(lista, solicitud)
	}

	return lista
}

func (s *SolicitudTutoriaStorage) GetByID(
	id int,
) (models.SolicitudTutoria, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	solicitud, exists := s.solicitudes[id]

	if !exists {
		return models.SolicitudTutoria{},
			errors.New("solicitud de tutoría no encontrada")
	}

	return solicitud, nil
}

func (s *SolicitudTutoriaStorage) Update(
	id int,
	solicitud models.SolicitudTutoria,
) (models.SolicitudTutoria, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.solicitudes[id]

	if !exists {
		return models.SolicitudTutoria{},
			errors.New("solicitud de tutoría no encontrada")
	}

	solicitud.ID = id
	s.solicitudes[id] = solicitud

	return solicitud, nil
}

func (s *SolicitudTutoriaStorage) Delete(id int) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.solicitudes[id]

	if !exists {
		return errors.New("solicitud de tutoría no encontrada")
	}

	delete(s.solicitudes, id)

	return nil
}