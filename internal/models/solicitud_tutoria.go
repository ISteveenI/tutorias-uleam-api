package models

type SolicitudTutoria struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	EstudianteID     int    `json:"estudiante_id"`
	HorarioDocenteID int    `json:"horario_docente_id"`
	TipoTutoriaID    int    `json:"tipo_tutoria_id"`
	Tema             string `json:"tema"`
	FechaSolicitud   string `json:"fecha_solicitud"`
	Estado           string `json:"estado"`

	TipoTutoria TipoTutoria `gorm:"foreignKey:TipoTutoriaID" json:"tipo_tutoria,omitempty"`
}