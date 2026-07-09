package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type EstudianteStorage struct {
	mu          sync.Mutex
	estudiantes map[uint]models.Estudiante
	nextID      uint
}

func NewEstudianteStorage() *EstudianteStorage {
	return &EstudianteStorage{
		estudiantes: make(map[uint]models.Estudiante),
		nextID:      1,
	}
}

func (e *EstudianteStorage) Create(estudiante models.Estudiante) models.Estudiante {

	e.mu.Lock()
	defer e.mu.Unlock()

	estudiante.ID = e.nextID
	e.estudiantes[estudiante.ID] = estudiante
	e.nextID++

	return estudiante
}

func (e *EstudianteStorage) GetAll() []models.Estudiante {

	e.mu.Lock()
	defer e.mu.Unlock()

	lista := make([]models.Estudiante, 0, len(e.estudiantes))

	for _, estudiante := range e.estudiantes {
		lista = append(lista, estudiante)
	}

	return lista
}

func (e *EstudianteStorage) GetByID(id uint) (models.Estudiante, error) {

	e.mu.Lock()
	defer e.mu.Unlock()

	estudiante, exists := e.estudiantes[id]

	if !exists {
		return models.Estudiante{}, errors.New("estudiante no encontrado")
	}

	return estudiante, nil
}

func (e *EstudianteStorage) Update(id uint, estudiante models.Estudiante) (models.Estudiante, error) {

	e.mu.Lock()
	defer e.mu.Unlock()

	_, exists := e.estudiantes[id]

	if !exists {
		return models.Estudiante{}, errors.New("estudiante no encontrado")
	}

	estudiante.ID = id
	e.estudiantes[id] = estudiante

	return estudiante, nil
}

func (e *EstudianteStorage) Delete(id uint) error {

	e.mu.Lock()
	defer e.mu.Unlock()

	_, exists := e.estudiantes[id]

	if !exists {
		return errors.New("estudiante no encontrado")
	}

	delete(e.estudiantes, id)

	return nil
}