package repositories

import (
	"context"
	"errors"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

var ErrDocenteNoEncontrado = errors.New("docente no encontrado")

type DocenteRepository interface {
	Create(ctx context.Context, docente *models.Docente) error
	FindByID(ctx context.Context, id uint) (*models.Docente, error)
	FindAll(ctx context.Context) ([]models.Docente, error)
	Update(ctx context.Context, docente *models.Docente) error
	Delete(ctx context.Context, id uint) error
}

type GormDocenteRepository struct {
	db *gorm.DB
}

func NewGormDocenteRepository(db *gorm.DB) *GormDocenteRepository {
	return &GormDocenteRepository{
		db: db,
	}
}

func (r *GormDocenteRepository) Create(
	ctx context.Context,
	docente *models.Docente,
) error {

	return r.db.WithContext(ctx).
		Create(docente).
		Error
}

func (r *GormDocenteRepository) FindByID(
	ctx context.Context,
	id uint,
) (*models.Docente, error) {

	var docente models.Docente

	err := r.db.WithContext(ctx).
		First(&docente, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrDocenteNoEncontrado
	}

	if err != nil {
		return nil, err
	}

	return &docente, nil
}

func (r *GormDocenteRepository) FindAll(
	ctx context.Context,
) ([]models.Docente, error) {

	var docentes []models.Docente

	err := r.db.WithContext(ctx).
		Find(&docentes).
		Error

	return docentes, err
}

func (r *GormDocenteRepository) Update(
	ctx context.Context,
	docente *models.Docente,
) error {

	result := r.db.WithContext(ctx).
		Save(docente)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrDocenteNoEncontrado
	}

	return nil
}

func (r *GormDocenteRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	result := r.db.WithContext(ctx).
		Delete(&models.Docente{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrDocenteNoEncontrado
	}

	return nil
}
