package repositories

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func TestGormDocenteRepository_CreateAndFindAll(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error al abrir la base de datos: %v", err)
	}

	err = db.AutoMigrate(
		&models.Docente{},
	)
	
	if err != nil {
		t.Fatalf("error al migrar: %v", err)
	}

	repo := NewGormDocenteRepository(db)

	docente := &models.Docente{
		Nombres:         "Karen",
		Apellidos:       "Holguin",
		Correo:          "karen@uleam.edu.ec",
		Telefono:        "0999999999",
		Departamento:    "Tecnologías",
		Especialidad:    "Aplicaciones Web",
		TituloAcademico: "Ingeniera",
	}

	err = repo.Create(context.Background(), docente)
	if err != nil {
		t.Fatalf("error al crear docente: %v", err)
	}

	docentes, err := repo.FindAll(context.Background())
	if err != nil {
		t.Fatalf("error al listar docentes: %v", err)
	}

	if len(docentes) != 1 {
		t.Fatalf("se esperaba 1 docente y se obtuvo %d", len(docentes))
	}

	if docentes[0].Correo != "karen@uleam.edu.ec" {
		t.Fatalf("correo incorrecto")
	}
}