package repositories

import (
	"context"
	"errors"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

type SolicitudTutoriaRepository interface {
	Create(ctx context.Context, solicitud *models.SolicitudTutoria) error
	FindByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error)
	FindAll(ctx context.Context) ([]models.SolicitudTutoria, error)
	Update(ctx context.Context, solicitud *models.SolicitudTutoria) error
	Delete(ctx context.Context, id uint) error
}

type GormSolicitudTutoriaRepository struct {
	db *gorm.DB
}

func NewGormSolicitudTutoriaRepository(db *gorm.DB) *GormSolicitudTutoriaRepository {
	return &GormSolicitudTutoriaRepository{
		db: db,
	}
}

func (r *GormSolicitudTutoriaRepository) Create(ctx context.Context, solicitud *models.SolicitudTutoria) error {
	return r.db.WithContext(ctx).Create(solicitud).Error
}

func (r *GormSolicitudTutoriaRepository) FindByID(ctx context.Context, id uint) (*models.SolicitudTutoria, error) {

	var solicitud models.SolicitudTutoria

	err := r.db.WithContext(ctx).First(&solicitud, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRegistroNoEncontrado
	}

	if err != nil {
		return nil, err
	}

	return &solicitud, nil
}

func (r *GormSolicitudTutoriaRepository) FindAll(ctx context.Context) ([]models.SolicitudTutoria, error) {

	var solicitudes []models.SolicitudTutoria

	err := r.db.WithContext(ctx).Find(&solicitudes).Error

	return solicitudes, err
}

func (r *GormSolicitudTutoriaRepository) Update(ctx context.Context, solicitud *models.SolicitudTutoria) error {

	var existente models.SolicitudTutoria

	err := r.db.WithContext(ctx).First(&existente, solicitud.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRegistroNoEncontrado
	}

	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Save(solicitud).Error
}

func (r *GormSolicitudTutoriaRepository) Delete(ctx context.Context, id uint) error {

	result := r.db.WithContext(ctx).Delete(&models.SolicitudTutoria{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrRegistroNoEncontrado
	}

	return nil
}