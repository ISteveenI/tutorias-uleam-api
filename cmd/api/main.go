package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/database"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

func main() {

	db, err := database.ConnectSQLite("tutorias.db")
	if err != nil {
		fmt.Println("Error al conectar la base de datos:", err)
		return
	}

	// Sesiones
	sesionRepo := repositories.NewGormSesionRepository(db)
	sesionService := services.NewSesionService(sesionRepo)
	sesionHandler := handlers.NewSesionHandler(sesionService)

	// Docentes
	docenteRepo := repositories.NewGormDocenteRepository(db)
	docenteService := services.NewDocenteService(docenteRepo)
	docenteHandler := handlers.NewDocenteHandler(docenteService)

	// Solicitudes
	solicitudRepo := repositories.NewGormSolicitudTutoriaRepository(db)
	solicitudService := services.NewSolicitudTutoriaService(solicitudRepo)
	solicitudHandler := handlers.NewSolicitudTutoriaHandler(solicitudService)

	r := chi.NewRouter()

	// Storage de módulos antiguos
	storage.NewDocenteStorage()
	storage.NewMateriaStorage()
	storage.NewHorarioStorage()

	// Rutas Docentes
	r.Mount("/api/v1/docentes", docenteHandler.Routes())

	// Rutas Materias
	r.Post("/api/v1/materias", handlers.CreateMateria)
	r.Get("/api/v1/materias", handlers.GetMaterias)
	r.Get("/api/v1/materias/{id}", handlers.GetMateriaByID)
	r.Put("/api/v1/materias/{id}", handlers.UpdateMateria)
	r.Delete("/api/v1/materias/{id}", handlers.DeleteMateria)

	// Rutas Horarios Docente
	r.Post("/api/v1/horarios-docente", handlers.CreateHorario)
	r.Get("/api/v1/horarios-docente", handlers.GetHorarios)
	r.Get("/api/v1/horarios-docente/{id}", handlers.GetHorarioByID)
	r.Put("/api/v1/horarios-docente/{id}", handlers.UpdateHorario)
	r.Delete("/api/v1/horarios-docente/{id}", handlers.DeleteHorario)
	
	// Rutas Estudiantes
	r.Post("/api/v1/estudiantes", handlers.CreateEstudiante)
	r.Get("/api/v1/estudiantes", handlers.GetEstudiantes)
	r.Get("/api/v1/estudiantes/{id}", handlers.GetEstudianteByID)
	r.Put("/api/v1/estudiantes/{id}", handlers.UpdateEstudiante)
	r.Delete("/api/v1/estudiantes/{id}", handlers.DeleteEstudiante)

	// Tipos de Tutoría
	r.Post("/api/v1/tipos-tutoria", handlers.CreateTipoTutoria)
	r.Get("/api/v1/tipos-tutoria", handlers.GetTiposTutoria)
	r.Get("/api/v1/tipos-tutoria/{id}", handlers.GetTipoTutoriaByID)
	r.Put("/api/v1/tipos-tutoria/{id}", handlers.UpdateTipoTutoria)
	r.Delete("/api/v1/tipos-tutoria/{id}", handlers.DeleteTipoTutoria)

	// Solicitudes de Tutoría
	r.Post("/api/v1/solicitudes-tutoria", handlers.CreateSolicitudTutoria)
	r.Get("/api/v1/solicitudes-tutoria", handlers.GetSolicitudesTutoria)
	r.Get("/api/v1/solicitudes-tutoria/{id}", handlers.GetSolicitudByID)
	r.Put("/api/v1/solicitudes-tutoria/{id}", handlers.UpdateSolicitudTutoria)
	r.Delete("/api/v1/solicitudes-tutoria/{id}", handlers.DeleteSolicitudTutoria)

	// Rutas Sesiones de Tutoría
	r.Mount("/api/v1/sesiones-tutoria", sesionHandler.Routes())

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}