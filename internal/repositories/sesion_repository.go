package repositories

import (
	"context"
	"errors"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

var ErrRegistroNoEncontrado = errors.New("registro no encontrado")

type SesionRepository interface {
	Create(ctx context.Context, sesion *models.SesionTutoria) error
	FindByID(ctx context.Context, id uint) (*models.SesionTutoria, error)
	FindAll(ctx context.Context) ([]models.SesionTutoria, error)
	Update(ctx context.Context, sesion *models.SesionTutoria) error
	Delete(ctx context.Context, id uint) error

	CreateAsistencia(ctx context.Context, asistencia *models.Asistencia) error
	CreateEvidencia(ctx context.Context, evidencia *models.Evidencia) error
}

type GormSesionRepository struct {
	db *gorm.DB
}

func NewGormSesionRepository(db *gorm.DB) *GormSesionRepository {
	return &GormSesionRepository{db: db}
}

func (r *GormSesionRepository) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	return r.db.WithContext(ctx).Create(sesion).Error
}

func (r *GormSesionRepository) FindByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	var sesion models.SesionTutoria

	err := r.db.WithContext(ctx).
		Preload("Asistencias").
		Preload("Evidencias").
		First(&sesion, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRegistroNoEncontrado
	}

	if err != nil {
		return nil, err
	}

	return &sesion, nil
}

func (r *GormSesionRepository) FindAll(ctx context.Context) ([]models.SesionTutoria, error) {
	var sesiones []models.SesionTutoria

	err := r.db.WithContext(ctx).
		Preload("Asistencias").
		Preload("Evidencias").
		Find(&sesiones).Error

	return sesiones, err
}

func (r *GormSesionRepository) Update(ctx context.Context, sesion *models.SesionTutoria) error {
	result := r.db.WithContext(ctx).Save(sesion)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrRegistroNoEncontrado
	}

	return nil
}

func (r *GormSesionRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.SesionTutoria{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrRegistroNoEncontrado
	}

	return nil
}

func (r *GormSesionRepository) CreateAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	return r.db.WithContext(ctx).Create(asistencia).Error
}

func (r *GormSesionRepository) CreateEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	return r.db.WithContext(ctx).Create(evidencia).Error
}
