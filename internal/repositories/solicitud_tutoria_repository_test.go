package repositories

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func setupSolicitudRepository(t *testing.T) *GormSolicitudTutoriaRepository {

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

	return NewGormSolicitudTutoriaRepository(db)
}

func crearSolicitudPrueba() *models.SolicitudTutoria {
	return &models.SolicitudTutoria{
		EstudianteID:     1,
		HorarioDocenteID: 2,
		TipoTutoriaID:    1,
		Tema:             "POO",
		FechaSolicitud:   "2026-06-30",
		Estado:           "Pendiente",
	}
}

func TestSolicitudRepository_CreateAndFindAll(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()

	err := repo.Create(context.Background(), solicitud)
	if err != nil {
		t.Fatalf("error al crear: %v", err)
	}

	lista, err := repo.FindAll(context.Background())
	if err != nil {
		t.Fatalf("error al listar: %v", err)
	}

	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 registro y se obtuvo %d", len(lista))
	}

	if lista[0].Tema != "POO" {
		t.Fatal("tema incorrecto")
	}
}

func TestSolicitudRepository_FindByID(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()

	err := repo.Create(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	encontrada, err := repo.FindByID(context.Background(), solicitud.ID)
	if err != nil {
		t.Fatal(err)
	}

	if encontrada.ID != solicitud.ID {
		t.Fatal("id incorrecto")
	}
}

func TestSolicitudRepository_Update(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()

	err := repo.Create(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	solicitud.Tema = "Bases de Datos"
	solicitud.Estado = "Aprobada"

	err = repo.Update(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	actualizada, err := repo.FindByID(context.Background(), solicitud.ID)
	if err != nil {
		t.Fatal(err)
	}

	if actualizada.Tema != "Bases de Datos" {
		t.Fatal("no se actualizó el tema")
	}

	if actualizada.Estado != "Aprobada" {
		t.Fatal("no se actualizó el estado")
	}
}

func TestSolicitudRepository_Delete(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()

	err := repo.Create(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete(context.Background(), solicitud.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.FindByID(context.Background(), solicitud.ID)

	if err == nil {
		t.Fatal("el registro debió eliminarse")
	}
}

func TestSolicitudRepository_FindByID_NotFound(t *testing.T) {

	repo := setupSolicitudRepository(t)

	_, err := repo.FindByID(context.Background(), 999)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}

func TestSolicitudRepository_Delete_NotFound(t *testing.T) {

	repo := setupSolicitudRepository(t)

	err := repo.Delete(context.Background(), 999)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}

func TestSolicitudRepository_Update_NotFound(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()
	solicitud.ID = 999

	err := repo.Update(context.Background(), solicitud)

	if err == nil {
		t.Fatal("se esperaba un error")
	}
}

func TestSolicitudRepository_CreateMultipleAndFindAll(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud1 := crearSolicitudPrueba()

	solicitud2 := crearSolicitudPrueba()
	solicitud2.Tema = "Bases de Datos"

	err := repo.Create(context.Background(), solicitud1)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Create(context.Background(), solicitud2)
	if err != nil {
		t.Fatal(err)
	}

	lista, err := repo.FindAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(lista) != 2 {
		t.Fatalf("se esperaban 2 registros y se obtuvieron %d", len(lista))
	}
}

func TestSolicitudRepository_UpdatePersistsChanges(t *testing.T) {

	repo := setupSolicitudRepository(t)

	solicitud := crearSolicitudPrueba()

	err := repo.Create(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	solicitud.Tema = "Ingeniería de Software"
	solicitud.Estado = "Aceptada"

	err = repo.Update(context.Background(), solicitud)
	if err != nil {
		t.Fatal(err)
	}

	encontrada, err := repo.FindByID(context.Background(), solicitud.ID)
	if err != nil {
		t.Fatal(err)
	}

	if encontrada.Tema != "Ingeniería de Software" {
		t.Fatal("el tema no se actualizó")
	}

	if encontrada.Estado != "Aceptada" {
		t.Fatal("el estado no se actualizó")
	}
}
