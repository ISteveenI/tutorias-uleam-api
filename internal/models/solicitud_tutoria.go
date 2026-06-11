package models

type SolicitudTutoria struct {
	ID               int    `json:"id"`
	EstudianteID     int    `json:"estudiante_id"`
	DisponibilidadID int    `json:"disponibilidad_id"`
	Motivo           string `json:"motivo"`
	Estado           string `json:"estado"`
}
