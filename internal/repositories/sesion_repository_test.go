package repositories

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func nuevaDBSesionTest(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(
		&models.SesionTutoria{},
		&models.Asistencia{},
		&models.Evidencia{},
	)
	require.NoError(t, err)

	return db
}

func nuevaSesionRepositoryTest() *models.SesionTutoria {
	return &models.SesionTutoria{
		SolicitudID:   1,
		FechaSesion:   "2026-07-01",
		HoraInicio:    "09:00",
		HoraFin:       "10:00",
		Observaciones: "Sesion de prueba",
		Estado:        "Programada",
	}
}

func TestGormSesionRepositoryCreateFindByIDYFindAll(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion := nuevaSesionRepositoryTest()

	err := repo.Create(ctx, sesion)
	require.NoError(t, err)
	require.NotZero(t, sesion.ID)

	encontrada, err := repo.FindByID(ctx, sesion.ID)
	require.NoError(t, err)
	require.NotNil(t, encontrada)

	assert.Equal(t, sesion.ID, encontrada.ID)
	assert.Equal(t, uint(1), encontrada.SolicitudID)
	assert.Equal(t, "Programada", encontrada.Estado)

	sesiones, err := repo.FindAll(ctx)
	require.NoError(t, err)

	assert.Len(t, sesiones, 1)
	assert.Equal(t, sesion.ID, sesiones[0].ID)
}

func TestGormSesionRepositoryFindByIDNoEncontrado(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion, err := repo.FindByID(ctx, 999)

	require.ErrorIs(t, err, ErrRegistroNoEncontrado)
	assert.Nil(t, sesion)
}

func TestGormSesionRepositoryUpdate(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion := nuevaSesionRepositoryTest()

	err := repo.Create(ctx, sesion)
	require.NoError(t, err)

	sesion.Estado = "Finalizada"
	sesion.Observaciones = "Sesion actualizada"

	err = repo.Update(ctx, sesion)
	require.NoError(t, err)

	actualizada, err := repo.FindByID(ctx, sesion.ID)
	require.NoError(t, err)

	assert.Equal(t, "Finalizada", actualizada.Estado)
	assert.Equal(t, "Sesion actualizada", actualizada.Observaciones)
}

func TestGormSesionRepositoryDelete(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion := nuevaSesionRepositoryTest()

	err := repo.Create(ctx, sesion)
	require.NoError(t, err)

	err = repo.Delete(ctx, sesion.ID)
	require.NoError(t, err)

	eliminada, err := repo.FindByID(ctx, sesion.ID)

	require.ErrorIs(t, err, ErrRegistroNoEncontrado)
	assert.Nil(t, eliminada)
}

func TestGormSesionRepositoryDeleteNoEncontrado(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	err := repo.Delete(ctx, 999)

	require.ErrorIs(t, err, ErrRegistroNoEncontrado)
}

func TestGormSesionRepositoryCreateAsistencia(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion := nuevaSesionRepositoryTest()
	err := repo.Create(ctx, sesion)
	require.NoError(t, err)

	asistencia := &models.Asistencia{
		SesionID:          sesion.ID,
		EstudianteAsistio: true,
		DocenteAsistio:    true,
		Observacion:       "Asistieron ambos",
	}

	err = repo.CreateAsistencia(ctx, asistencia)
	require.NoError(t, err)
	require.NotZero(t, asistencia.ID)

	encontrada, err := repo.FindByID(ctx, sesion.ID)
	require.NoError(t, err)

	require.Len(t, encontrada.Asistencias, 1)
	assert.Equal(t, asistencia.ID, encontrada.Asistencias[0].ID)
	assert.True(t, encontrada.Asistencias[0].EstudianteAsistio)
	assert.True(t, encontrada.Asistencias[0].DocenteAsistio)
}

func TestGormSesionRepositoryCreateEvidencia(t *testing.T) {
	db := nuevaDBSesionTest(t)
	repo := NewGormSesionRepository(db)
	ctx := context.Background()

	sesion := nuevaSesionRepositoryTest()
	err := repo.Create(ctx, sesion)
	require.NoError(t, err)

	evidencia := &models.Evidencia{
		SesionID:    sesion.ID,
		TipoArchivo: "pdf",
		ArchivoURL:  "https://example.com/evidencia.pdf",
		Descripcion: "Evidencia de prueba",
	}

	err = repo.CreateEvidencia(ctx, evidencia)
	require.NoError(t, err)
	require.NotZero(t, evidencia.ID)

	encontrada, err := repo.FindByID(ctx, sesion.ID)
	require.NoError(t, err)

	require.Len(t, encontrada.Evidencias, 1)
	assert.Equal(t, evidencia.ID, encontrada.Evidencias[0].ID)
	assert.Equal(t, "pdf", encontrada.Evidencias[0].TipoArchivo)
	assert.Equal(t, "https://example.com/evidencia.pdf", encontrada.Evidencias[0].ArchivoURL)
}
