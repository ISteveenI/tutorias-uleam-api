package database

import (
	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func ConnectSQLite(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(

		&models.Docente{},
		&models.SesionTutoria{},
		&models.Asistencia{},
		&models.Evidencia{},

		// Modulo Solicitudes
		&models.TipoTutoria{},
		&models.SolicitudTutoria{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
