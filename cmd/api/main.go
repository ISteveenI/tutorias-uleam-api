package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/database"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/seeders"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

func main() {

	db, err := database.ConnectPostgres()
	if err != nil {
		fmt.Println("Error al conectar la base de datos:", err)
		return
	}

	seeders.SeedUsers(db)

	// ==========================
	// AUTH
	// ==========================

	userRepo := repositories.NewGormUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// ==========================
	// SESIONES
	// ==========================

	sesionRepo := repositories.NewGormSesionRepository(db)
	sesionService := services.NewSesionService(sesionRepo)
	sesionHandler := handlers.NewSesionHandler(sesionService)

	// ==========================
	// DOCENTES
	// ==========================

	docenteRepo := repositories.NewGormDocenteRepository(db)
	docenteService := services.NewDocenteService(docenteRepo)
	docenteHandler := handlers.NewDocenteHandler(docenteService)

	// ==========================
	// SOLICITUDES
	// ==========================

	solicitudRepo := repositories.NewGormSolicitudTutoriaRepository(db)
	solicitudService := services.NewSolicitudTutoriaService(solicitudRepo)
	solicitudHandler := handlers.NewSolicitudTutoriaHandler(solicitudService)

	r := chi.NewRouter()

	// ==========================
	// AUTH
	// ==========================

	r.Post("/api/v1/auth/register", authHandler.Register)
	r.Post("/api/v1/auth/login", authHandler.Login)

	// ==========================
	// STORAGE
	// ==========================

	storage.NewDocenteStorage()
	storage.NewMateriaStorage()
	storage.NewHorarioStorage()

	// ==========================
	// DOCENTES (ADMIN)
	// ==========================

	r.Route("/api/v1/docentes", func(r chi.Router) {

		r.Use(middlewares.RequireAuth)
		r.Use(middlewares.RequireRole("admin"))

		r.Mount("/", docenteHandler.Routes())
	})

	// ==========================
	// MATERIAS
	// ==========================

	r.Post("/api/v1/materias", handlers.CreateMateria)
	r.Get("/api/v1/materias", handlers.GetMaterias)
	r.Get("/api/v1/materias/{id}", handlers.GetMateriaByID)
	r.Put("/api/v1/materias/{id}", handlers.UpdateMateria)
	r.Delete("/api/v1/materias/{id}", handlers.DeleteMateria)

	// ==========================
	// HORARIOS
	// ==========================

	r.Post("/api/v1/horarios-docente", handlers.CreateHorario)
	r.Get("/api/v1/horarios-docente", handlers.GetHorarios)
	r.Get("/api/v1/horarios-docente/{id}", handlers.GetHorarioByID)
	r.Put("/api/v1/horarios-docente/{id}", handlers.UpdateHorario)
	r.Delete("/api/v1/horarios-docente/{id}", handlers.DeleteHorario)

	// ==========================
	// ESTUDIANTES
	// ==========================

	r.Post("/api/v1/estudiantes", handlers.CreateEstudiante)
	r.Get("/api/v1/estudiantes", handlers.GetEstudiantes)
	r.Get("/api/v1/estudiantes/{id}", handlers.GetEstudianteByID)
	r.Put("/api/v1/estudiantes/{id}", handlers.UpdateEstudiante)
	r.Delete("/api/v1/estudiantes/{id}", handlers.DeleteEstudiante)

	// ==========================
	// TIPOS DE TUTORÍA
	// ==========================

	r.Post("/api/v1/tipos-tutoria", handlers.CreateTipoTutoria)
	r.Get("/api/v1/tipos-tutoria", handlers.GetTiposTutoria)
	r.Get("/api/v1/tipos-tutoria/{id}", handlers.GetTipoTutoriaByID)
	r.Put("/api/v1/tipos-tutoria/{id}", handlers.UpdateTipoTutoria)
	r.Delete("/api/v1/tipos-tutoria/{id}", handlers.DeleteTipoTutoria)

	// ==========================
	// SOLICITUDES (USUARIOS AUTENTICADOS)
	// ==========================

	r.Route("/api/v1/solicitudes-tutoria", func(r chi.Router) {

		r.Use(middlewares.RequireAuth)

		r.Mount("/", solicitudHandler.Routes())
	})

	// ==========================
	// SESIONES (DOCENTE Y ADMIN)
	// ==========================

	r.Route("/api/v1/sesiones-tutoria", func(r chi.Router) {

		r.Use(middlewares.RequireAuth)
		r.Use(middlewares.RequireRole("docente", "admin"))

		r.Mount("/", sesionHandler.Routes())
	})

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
