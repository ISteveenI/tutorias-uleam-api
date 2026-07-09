package database

import (
	"time"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDatosIniciales(db *gorm.DB) error {
	if err := seedUsuarios(db); err != nil {
		return err
	}

	if err := seedSesiones(db); err != nil {
		return err
	}

	return nil
}

func seedUsuarios(db *gorm.DB) error {
	var totalUsuarios int64

	if err := db.Model(&models.Usuario{}).Count(&totalUsuarios).Error; err != nil {
		return err
	}

	if totalUsuarios > 0 {
		return nil
	}

	usuarios := []struct {
		Nombre   string
		Correo   string
		Password string
		Rol      string
	}{
		{
			Nombre:   "Administrador",
			Correo:   "admin@uleam.edu.ec",
			Password: "admin123",
			Rol:      "admin",
		},
		{
			Nombre:   "Docente Demo",
			Correo:   "docente@uleam.edu.ec",
			Password: "docente123",
			Rol:      "docente",
		},
		{
			Nombre:   "Estudiante Demo",
			Correo:   "estudiante@uleam.edu.ec",
			Password: "estudiante123",
			Rol:      "estudiante",
		},
	}

	for _, item := range usuarios {
		hash, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		usuario := models.Usuario{
			Nombre:       item.Nombre,
			Correo:       item.Correo,
			PasswordHash: string(hash),
			Rol:          item.Rol,
		}

		if err := db.Create(&usuario).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedSesiones(db *gorm.DB) error {
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
