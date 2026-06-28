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

	// Rutas Disponibilidad Docente
	


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
