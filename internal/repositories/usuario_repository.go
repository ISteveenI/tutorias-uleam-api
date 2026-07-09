package repositories

import (
	"context"
	"errors"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	Create(ctx context.Context, usuario *models.Usuario) error
	FindByCorreo(ctx context.Context, correo string) (*models.Usuario, error)
}

type GormUsuarioRepository struct {
	db *gorm.DB
}

func NewGormUsuarioRepository(db *gorm.DB) *GormUsuarioRepository {
	return &GormUsuarioRepository{db: db}
}

func (r *GormUsuarioRepository) Create(ctx context.Context, usuario *models.Usuario) error {
	return r.db.WithContext(ctx).Create(usuario).Error
}

func (r *GormUsuarioRepository) FindByCorreo(ctx context.Context, correo string) (*models.Usuario, error) {
	var usuario models.Usuario

	err := r.db.WithContext(ctx).
		Where("correo = ?", correo).
		First(&usuario).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRegistroNoEncontrado
	}

	if err != nil {
		return nil, err
	}

	return &usuario, nil
}
