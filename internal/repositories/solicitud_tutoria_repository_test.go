package repositories

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func TestSolicitudRepository_CreateAndFindAll(t *testing.T) {

	// Configurar la base de datos en memoria para pruebas
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite: %v", err)
	}

	// Migrar los modelos necesarios para la prueba
	err = db.AutoMigrate(
		&models.TipoTutoria{},
		&models.SolicitudTutoria{},
	)

	if err != nil {
		t.Fatalf("error automigrate: %v", err)
	}

	// Crear el repositorio
	repo := NewGormSolicitudTutoriaRepository(db)

	// Crear una solicitud de prueba
	solicitud := &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 2,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
		Estado:           "Pendiente",
	}

	// Guardar la solicitud en la base de datos
	err = repo.Create(context.Background(), solicitud)

	if err != nil {
		t.Fatalf("error al crear: %v", err)
	}

	// Listar todas las solicitudes
	solicitudes, err := repo.FindAll(context.Background())

	if err != nil {
		t.Fatalf("error al listar: %v", err)
	}

	// Verificar que se haya creado correctamente
	if len(solicitudes) != 1 {
		t.Fatalf("se esperaba 1 solicitud y se obtuvo %d", len(solicitudes))
	}

	// Verificar que los datos sean correctos
	if solicitudes[0].Tema != "POO" {
		t.Fatalf("tema incorrecto")
	}
}
