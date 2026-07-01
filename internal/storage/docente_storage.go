package storage

import (
	"errors"
	"sync"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

type DocenteStorage struct {
	mu        sync.Mutex
	docentes  map[int]models.Docente
	nextID    int
}

func NewDocenteStorage() *DocenteStorage {
	return &DocenteStorage{
		docentes: make(map[int]models.Docente),
		nextID:   1,
	}
}

func (d *DocenteStorage) Create(docente models.Docente) models.Docente {

	d.mu.Lock()
	defer d.mu.Unlock()

	docente.ID = uint(d.nextID)
	d.docentes[d.nextID] = docente
	d.nextID++
	return docente
}

func (d *DocenteStorage) GetAll() []models.Docente {
	d.mu.Lock()
	defer d.mu.Unlock()
	lista := make([]models.Docente, 0, len(d.docentes))
	for _, docente := range d.docentes {
		lista = append(lista, docente)
	}
	return lista
}

func (d *DocenteStorage) GetByID(id int) (models.Docente, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	docente, exists := d.docentes[id]
	if !exists {
		return models.Docente{}, errors.New("docente no encontrado")
	}
	return docente, nil
}

func (d *DocenteStorage) Update(id int,docente models.Docente,) (models.Docente, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, exists := d.docentes[id]
	if !exists {
		return models.Docente{}, errors.New("docente no encontrado")
	}
	docente.ID = uint(id)
	d.docentes[id] = docente
	return docente, nil
}

func (d *DocenteStorage) Delete(id int) error {

	d.mu.Lock()
	defer d.mu.Unlock()
	_, exists := d.docentes[id]
	if !exists {
		return errors.New("docente no encontrado")
	}
	delete(d.docentes, id)
	return nil
}

var DocenteRepo = NewDocenteStorage()