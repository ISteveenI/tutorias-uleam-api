package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type MateriaStorage struct {
	mu        sync.Mutex
	materias  map[int]models.Materia
	nextID    int
}

func NewMateriaStorage() *MateriaStorage {
	return &MateriaStorage{
		materias: make(map[int]models.Materia),
		nextID:   1,
	}
}

func (m *MateriaStorage) Create(materia models.Materia) models.Materia {

	m.mu.Lock()
	defer m.mu.Unlock()
	materia.ID = m.nextID
	m.materias[materia.ID] = materia
	m.nextID++
	return materia
}

func (m *MateriaStorage) GetAll() []models.Materia {

	m.mu.Lock()
	defer m.mu.Unlock()
	lista := make([]models.Materia, 0, len(m.materias))
	for _, materia := range m.materias {
		lista = append(lista, materia)
	}
	return lista
}

func (m *MateriaStorage) GetByID(id int) (models.Materia, error) {

	m.mu.Lock()
	defer m.mu.Unlock()
	materia, exists := m.materias[id]
	if !exists {
		return models.Materia{}, errors.New("materia no encontrada")
	}
	return materia, nil
}

func (m *MateriaStorage) Update(id int,materia models.Materia,) (models.Materia, error) {

	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.materias[id]
	if !exists {
		return models.Materia{}, errors.New("materia no encontrada")
	}
	materia.ID = id
	m.materias[id] = materia
	return materia, nil
}

func (m *MateriaStorage) Delete(id int) error {

	m.mu.Lock()
	defer m.mu.Unlock()
	_, exists := m.materias[id]
	if !exists {
		return errors.New("materia no encontrada")
	}
	delete(m.materias, id)
	return nil
}

var MateriaRepo = NewMateriaStorage()