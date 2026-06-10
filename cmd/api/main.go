package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/handlers"
)

func main() {

	r := chi.NewRouter()
	// Rutas para Disponibilidades
	r.Post("/api/v1/disponibilidades", handlers.CreateDisponibilidad)
	r.Get("/api/v1/disponibilidades", handlers.GetDisponibilidades)
	r.Get("/api/v1/disponibilidades/{id}", handlers.GetDisponibilidadByID)
	r.Put("/api/v1/disponibilidades/{id}", handlers.UpdateDisponibilidad)
	r.Delete("/api/v1/disponibilidades/{id}", handlers.DeleteDisponibilidad)

	// Rutas para Solicitudes de Tutoría
	r.Post("/api/v1/solicitudes-tutoria", handlers.CreateSolicitudTutoria)
	r.Get("/api/v1/solicitudes-tutoria", handlers.SolicitudesTutoria)

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
