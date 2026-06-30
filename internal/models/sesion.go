package models

import "time"

type SesionTutoria struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	SolicitudID   uint   `gorm:"not null;index" json:"solicitud_id"`
	FechaSesion   string `gorm:"not null" json:"fecha_sesion"`
	HoraInicio    string `gorm:"not null" json:"hora_inicio"`
	HoraFin       string `gorm:"not null" json:"hora_fin"`
	Observaciones string `json:"observaciones"`
	Estado        string `gorm:"not null" json:"estado"`

	Asistencias []Asistencia `gorm:"foreignKey:SesionID" json:"asistencias,omitempty"`
	Evidencias  []Evidencia  `gorm:"foreignKey:SesionID" json:"evidencias,omitempty"`
}

type Asistencia struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	SesionID          uint   `gorm:"not null;index" json:"sesion_id"`
	EstudianteAsistio bool   `json:"estudiante_asistio"`
	DocenteAsistio    bool   `json:"docente_asistio"`
	Observacion       string `json:"observacion"`
}

type Evidencia struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	SesionID    uint      `gorm:"not null;index" json:"sesion_id"`
	TipoArchivo string    `gorm:"not null" json:"tipo_archivo"`
	ArchivoURL  string    `gorm:"not null" json:"archivo_url"`
	FechaSubida time.Time `gorm:"autoCreateTime" json:"fecha_subida"`
	Descripcion string    `json:"descripcion"`
}
