package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

func main() {

	r := chi.NewRouter()
	sesionStorage := storage.NewSesionStorage()
	sesionHandler := handlers.NewSesionHandler(sesionStorage)

		// Inicialización storage
	storage.NewDocenteStorage()
	storage.NewMateriaStorage()
	storage.NewHorarioStorage()

	// Rutas Docentes
	r.Post("/api/v1/docentes",handlers.CreateDocente,)
	r.Get("/api/v1/docentes",handlers.GetDocentes,)
	r.Get("/api/v1/docentes/{id}",handlers.GetDocenteByID,)
	r.Put("/api/v1/docentes/{id}",handlers.UpdateDocente,)
	r.Delete("/api/v1/docentes/{id}",handlers.DeleteDocente,)

	// Rutas Materias
	r.Post("/api/v1/materias",handlers.CreateMateria,)
	r.Get("/api/v1/materias",handlers.GetMaterias,)
	r.Get("/api/v1/materias/{id}",handlers.GetMateriaByID,)
	r.Put("/api/v1/materias/{id}",handlers.UpdateMateria,)
	r.Delete("/api/v1/materias/{id}",handlers.DeleteMateria,)

	// Rutas Horarios Docente
	r.Post("/api/v1/horarios-docente",handlers.CreateHorario,)
	r.Get("/api/v1/horarios-docente",handlers.GetHorarios,)
	r.Get("/api/v1/horarios-docente/{id}",handlers.GetHorarioByID,)
	r.Put("/api/v1/horarios-docente/{id}",handlers.UpdateHorario,)
	r.Delete("/api/v1/horarios-docente/{id}",handlers.DeleteHorario,)

	// Rutas Solicitudes de Tutoría
	r.Post("/api/v1/solicitudes-tutoria", handlers.CreateSolicitudTutoria)
	r.Get("/api/v1/solicitudes-tutoria", handlers.GetSolicitudesTutoria)
	r.Get("/api/v1/solicitudes/{id}", handlers.GetSolicitudByID)
	r.Put("/api/v1/solicitudes/{id}", handlers.UpdateSolicitudTutoria)
	r.Delete("/api/v1/solicitudes/{id}", handlers.DeleteSolicitudTutoria)

	// Rutas Sesiones de Tutoría
	r.Mount("/api/v1/sesiones-tutoria", sesionHandler.Routes())

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
