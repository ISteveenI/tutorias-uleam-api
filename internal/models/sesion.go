package models

type SesionTutoria struct {
	ID               int    `json:"id"`
	SolicitudID      int    `json:"solicitud_id"`
	DisponibilidadID int    `json:"disponibilidad_id"`
	DocenteID        int    `json:"docente_id"`
	EstudianteID     int    `json:"estudiante_id"`
	Fecha            string `json:"fecha"`
	HoraInicio       string `json:"hora_inicio"`
	HoraFin          string `json:"hora_fin"`
	Estado           string `json:"estado"`
	Observacion      string `json:"observacion"`
}
