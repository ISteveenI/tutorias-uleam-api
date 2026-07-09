package database

import (
	"time"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/gorm"
)

func SeedDatosIniciales(db *gorm.DB) error {
	var totalSesiones int64

	if err := db.Model(&models.SesionTutoria{}).Count(&totalSesiones).Error; err != nil {
		return err
	}

	if totalSesiones > 0 {
		return nil
	}

	sesion := models.SesionTutoria{
		SolicitudID:   1,
		FechaSesion:   "2026-07-01",
		HoraInicio:    "09:00",
		HoraFin:       "10:00",
		Observaciones: "Sesion inicial cargada por seeder",
		Estado:        "Programada",
	}

	if err := db.Create(&sesion).Error; err != nil {
		return err
	}

	asistencia := models.Asistencia{
		SesionID:          sesion.ID,
		EstudianteAsistio: true,
		DocenteAsistio:    true,
		Observacion:       "Asistencia inicial generada por seeder",
	}

	if err := db.Create(&asistencia).Error; err != nil {
		return err
	}

	evidencia := models.Evidencia{
		SesionID:    sesion.ID,
		TipoArchivo: "pdf",
		ArchivoURL:  "https://example.com/evidencia-inicial.pdf",
		FechaSubida: time.Now(),
		Descripcion: "Evidencia inicial generada por seeder",
	}

	if err := db.Create(&evidencia).Error; err != nil {
		return err
	}

	return nil
}
