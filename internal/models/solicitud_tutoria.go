package models

type SolicitudTutoria struct {
	ID                 int    `json:"id"`
	EstudianteID       int    `json:"estudiante_id"`
	DocenteID          int    `json:"docente_id"`
	DisponibilidadID   int    `json:"disponibilidad_id"`
	Materia            string `json:"materia"`
	Tema               string `json:"tema"`
	Urgencia           string `json:"urgencia"`
	Modalidad          string `json:"modalidad"`
	FechaSolicitud     string `json:"fecha_solicitud"`
	PrioridadCalculada int    `json:"prioridad_calculada"`
	Estado             string `json:"estado"`
}
