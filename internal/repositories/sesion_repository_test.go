package repositories_test

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"gorm.io/gorm"
)

func TestGormSesionRepositoryCreateYBuscarListar(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite en memoria: %v", err)
	}

	err = db.AutoMigrate(
		&models.SesionTutoria{},
		&models.Asistencia{},
		&models.Evidencia{},
	)

	if err != nil {
		t.Fatalf("no se pudo ejecutar AutoMigrate: %v", err)
	}

	repo := repositories.NewGormSesionRepository(db)

	sesion := &models.SesionTutoria{
		SolicitudID:   1,
		FechaSesion:   "2026-07-01",
		HoraInicio:    "09:00",
		HoraFin:       "10:00",
		Observaciones: "Sesion creada desde test",
		Estado:        "Programada",
	}

	err = repo.Create(context.Background(), sesion)
	if err != nil {
		t.Fatalf("no se pudo crear la sesion: %v", err)
	}

	if sesion.ID == 0 {
		t.Fatal("se esperaba que GORM asigne un ID a la sesion")
	}

	encontrada, err := repo.FindByID(context.Background(), sesion.ID)
	if err != nil {
		t.Fatalf("no se pudo buscar la sesion creada: %v", err)
	}

	if encontrada.SolicitudID != sesion.SolicitudID {
		t.Fatalf("se esperaba SolicitudID %d, pero se obtuvo %d", sesion.SolicitudID, encontrada.SolicitudID)
	}

	lista, err := repo.FindAll(context.Background())
	if err != nil {
		t.Fatalf("no se pudo listar sesiones: %v", err)
	}

	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 sesion en la lista, pero se obtuvo %d", len(lista))
	}
}
