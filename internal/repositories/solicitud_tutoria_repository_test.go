package repositories

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func TestSolicitudRepository_CreateAndFindAll(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite: %v", err)
	}

	err = db.AutoMigrate(
		&models.TipoTutoria{},
		&models.SolicitudTutoria{},
	)

	if err != nil {
		t.Fatalf("error automigrate: %v", err)
	}

	repo := NewGormSolicitudTutoriaRepository(db)

	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 2,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
		Estado:           "Pendiente",
	}

	err = repo.Create(context.Background(), solicitud)

	if err != nil {
		t.Fatalf("error al crear: %v", err)
	}

	solicitudes, err := repo.FindAll(context.Background())

	if err != nil {
		t.Fatalf("error al listar: %v", err)
	}

	if len(solicitudes) != 1 {
		t.Fatalf("se esperaba 1 solicitud y se obtuvo %d", len(solicitudes))
	}

	if solicitudes[0].Tema != "POO" {
		t.Fatalf("tema incorrecto")
	}
}