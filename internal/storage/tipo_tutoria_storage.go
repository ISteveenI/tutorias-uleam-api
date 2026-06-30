package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type TipoTutoriaStorage struct {
	mu            sync.Mutex
	tiposTutoria  map[int]models.TipoTutoria
	nextID        int
}

func NewTipoTutoriaStorage() *TipoTutoriaStorage {
	return &TipoTutoriaStorage{
		tiposTutoria: make(map[int]models.TipoTutoria),
		nextID:       1,
	}
}

func (t *TipoTutoriaStorage) Create(tipo models.TipoTutoria) models.TipoTutoria {

	t.mu.Lock()
	defer t.mu.Unlock()

	tipo.ID = t.nextID
	t.tiposTutoria[tipo.ID] = tipo
	t.nextID++

	return tipo
}

func (t *TipoTutoriaStorage) GetAll() []models.TipoTutoria {

	t.mu.Lock()
	defer t.mu.Unlock()

	lista := make([]models.TipoTutoria, 0, len(t.tiposTutoria))

	for _, tipo := range t.tiposTutoria {
		lista = append(lista, tipo)
	}

	return lista
}

func (t *TipoTutoriaStorage) GetByID(id int) (models.TipoTutoria, error) {

	t.mu.Lock()
	defer t.mu.Unlock()

	tipo, exists := t.tiposTutoria[id]

	if !exists {
		return models.TipoTutoria{}, errors.New("tipo de tutoría no encontrado")
	}

	return tipo, nil
}

func (t *TipoTutoriaStorage) Update(
	id int,
	tipo models.TipoTutoria,
) (models.TipoTutoria, error) {

	t.mu.Lock()
	defer t.mu.Unlock()

	_, exists := t.tiposTutoria[id]

	if !exists {
		return models.TipoTutoria{}, errors.New("tipo de tutoría no encontrado")
	}

	tipo.ID = id
	t.tiposTutoria[id] = tipo

	return tipo, nil
}

func (t *TipoTutoriaStorage) Delete(id int) error {

	t.mu.Lock()
	defer t.mu.Unlock()

	_, exists := t.tiposTutoria[id]

	if !exists {
		return errors.New("tipo de tutoría no encontrado")
	}

	delete(t.tiposTutoria, id)

	return nil
}