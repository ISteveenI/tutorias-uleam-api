package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockSesionRepository struct {
	mock.Mock
}

func (m *mockSesionRepository) Create(ctx context.Context, sesion *models.SesionTutoria) error {
	args := m.Called(ctx, sesion)
	return args.Error(0)
}

func (m *mockSesionRepository) FindByID(ctx context.Context, id uint) (*models.SesionTutoria, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.SesionTutoria), args.Error(1)
}

func (m *mockSesionRepository) FindAll(ctx context.Context) ([]models.SesionTutoria, error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.SesionTutoria), args.Error(1)
}

func (m *mockSesionRepository) Update(ctx context.Context, sesion *models.SesionTutoria) error {
	args := m.Called(ctx, sesion)
	return args.Error(0)
}

func (m *mockSesionRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockSesionRepository) CreateAsistencia(ctx context.Context, asistencia *models.Asistencia) error {
	args := m.Called(ctx, asistencia)
	return args.Error(0)
}

func (m *mockSesionRepository) CreateEvidencia(ctx context.Context, evidencia *models.Evidencia) error {
	args := m.Called(ctx, evidencia)
	return args.Error(0)
}

func sesionValidaParaTest() *models.SesionTutoria {
	return &models.SesionTutoria{
		SolicitudID:   1,
		FechaSesion:   "2026-07-01",
		HoraInicio:    "09:00",
		HoraFin:       "10:00",
		Observaciones: "Sesion de prueba",
	}
}

func TestCreateSesionValidaUsaEstadoPorDefectoYRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion := sesionValidaParaTest()

	repo.On("Create", mock.Anything, mock.MatchedBy(func(s *models.SesionTutoria) bool {
		return s.SolicitudID == 1 &&
			s.Estado == "Programada" &&
			s.FechaSesion == "2026-07-01" &&
			s.HoraInicio == "09:00" &&
			s.HoraFin == "10:00"
	})).Return(nil).Once()

	err := service.Create(context.Background(), sesion)

	require.NoError(t, err)
	assert.Equal(t, "Programada", sesion.Estado)
	repo.AssertExpectations(t)
}

func TestCreateSesionConEstadoPorDefectoPersonalizado(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo, WithEstadoPorDefecto("Pendiente"))

	sesion := sesionValidaParaTest()

	repo.On("Create", mock.Anything, mock.MatchedBy(func(s *models.SesionTutoria) bool {
		return s.Estado == "Pendiente"
	})).Return(nil).Once()

	err := service.Create(context.Background(), sesion)

	require.NoError(t, err)
	assert.Equal(t, "Pendiente", sesion.Estado)
	repo.AssertExpectations(t)
}

func TestCreateSesionConSolicitudIDInvalidoNoLlegaAlRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion := sesionValidaParaTest()
	sesion.SolicitudID = 0

	err := service.Create(context.Background(), sesion)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestCreateSesionConHoraInvalidaRetornaError(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion := sesionValidaParaTest()
	sesion.HoraInicio = "11:00"
	sesion.HoraFin = "10:00"

	err := service.Create(context.Background(), sesion)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestGetByIDValidoConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesionEsperada := &models.SesionTutoria{
		ID:          8,
		SolicitudID: 2,
		Estado:      "Programada",
	}

	repo.On("FindByID", mock.Anything, uint(8)).
		Return(sesionEsperada, nil).
		Once()

	sesion, err := service.GetByID(context.Background(), 8)

	require.NoError(t, err)
	require.NotNil(t, sesion)
	assert.Equal(t, uint(8), sesion.ID)
	assert.Equal(t, uint(2), sesion.SolicitudID)
	repo.AssertExpectations(t)
}

func TestGetByIDCeroRetornaErrorYNoConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion, err := service.GetByID(context.Background(), 0)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	assert.Nil(t, sesion)
	repo.AssertNotCalled(t, "FindByID", mock.Anything, mock.Anything)
}

func TestGetAllConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesionesEsperadas := []models.SesionTutoria{
		{ID: 1, SolicitudID: 1, Estado: "Programada"},
		{ID: 2, SolicitudID: 2, Estado: "Finalizada"},
	}

	repo.On("FindAll", mock.Anything).
		Return(sesionesEsperadas, nil).
		Once()

	sesiones, err := service.GetAll(context.Background())

	require.NoError(t, err)
	assert.Len(t, sesiones, 2)
	repo.AssertExpectations(t)
}

func TestUpdateValidoAsignaIDYConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion := sesionValidaParaTest()

	repo.On("Update", mock.Anything, mock.MatchedBy(func(s *models.SesionTutoria) bool {
		return s.ID == 15 &&
			s.SolicitudID == 1 &&
			s.FechaSesion == "2026-07-01" &&
			s.HoraInicio == "09:00" &&
			s.HoraFin == "10:00"
	})).Return(nil).Once()

	err := service.Update(context.Background(), 15, sesion)

	require.NoError(t, err)
	assert.Equal(t, uint(15), sesion.ID)
	repo.AssertExpectations(t)
}

func TestUpdateConIDCeroRetornaErrorYNoConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	err := service.Update(context.Background(), 0, sesionValidaParaTest())

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
}

func TestDeleteValidoConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	repo.On("Delete", mock.Anything, uint(20)).
		Return(nil).
		Once()

	err := service.Delete(context.Background(), 20)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteConIDCeroRetornaErrorYNoConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	err := service.Delete(context.Background(), 0)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "Delete", mock.Anything, mock.Anything)
}

func TestRegistrarAsistenciaValidaConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	asistencia := &models.Asistencia{
		SesionID:          3,
		EstudianteAsistio: true,
		DocenteAsistio:    true,
		Observacion:       "Asistieron ambos",
	}

	repo.On("CreateAsistencia", mock.Anything, mock.MatchedBy(func(a *models.Asistencia) bool {
		return a.SesionID == 3 &&
			a.EstudianteAsistio &&
			a.DocenteAsistio
	})).Return(nil).Once()

	err := service.RegistrarAsistencia(context.Background(), asistencia)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestRegistrarAsistenciaSinSesionIDRetornaError(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	asistencia := &models.Asistencia{
		SesionID:          0,
		EstudianteAsistio: true,
		DocenteAsistio:    true,
	}

	err := service.RegistrarAsistencia(context.Background(), asistencia)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "CreateAsistencia", mock.Anything, mock.Anything)
}

func TestSubirEvidenciaValidaSeteaFechaYConsultaRepositorio(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	evidencia := &models.Evidencia{
		SesionID:    4,
		TipoArchivo: "pdf",
		ArchivoURL:  "https://example.com/evidencia.pdf",
		Descripcion: "Evidencia de la tutoria",
		FechaSubida: time.Time{},
	}

	repo.On("CreateEvidencia", mock.Anything, mock.MatchedBy(func(e *models.Evidencia) bool {
		return e.SesionID == 4 &&
			e.TipoArchivo == "pdf" &&
			e.ArchivoURL == "https://example.com/evidencia.pdf" &&
			!e.FechaSubida.IsZero()
	})).Return(nil).Once()

	err := service.SubirEvidencia(context.Background(), evidencia)

	require.NoError(t, err)
	assert.False(t, evidencia.FechaSubida.IsZero())
	assert.WithinDuration(t, time.Now(), evidencia.FechaSubida, time.Second)
	repo.AssertExpectations(t)
}

func TestSubirEvidenciaSinArchivoURLRetornaError(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	evidencia := &models.Evidencia{
		SesionID:    4,
		TipoArchivo: "pdf",
		ArchivoURL:  "",
		Descripcion: "Evidencia sin URL",
	}

	err := service.SubirEvidencia(context.Background(), evidencia)

	require.ErrorIs(t, err, ErrDatosInvalidos)
	repo.AssertNotCalled(t, "CreateEvidencia", mock.Anything, mock.Anything)
}

func TestRepositorioPropagaErrorEnCreate(t *testing.T) {
	repo := new(mockSesionRepository)
	service := NewSesionService(repo)

	sesion := sesionValidaParaTest()
	errorRepositorio := errors.New("error de base de datos")

	repo.On("Create", mock.Anything, mock.AnythingOfType("*models.SesionTutoria")).
		Return(errorRepositorio).
		Once()

	err := service.Create(context.Background(), sesion)

	require.ErrorIs(t, err, errorRepositorio)
	repo.AssertExpectations(t)
}
