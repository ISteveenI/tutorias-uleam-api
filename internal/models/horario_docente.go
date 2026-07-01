package models

type DisponibilidadDocente struct {
	ID         int    `json:"id"`
	DocenteID  int    `json:"docente_id"`
	Materia    string `json:"materia"`
	DiaSemana  string `json:"dia_semana"`
	HoraInicio string `json:"hora_inicio"`
	HoraFin    string `json:"hora_fin"`
	Modalidad  string `json:"modalidad"`
	Cupos      int    `json:"cupos"`
	Estado     string `json:"estado"`
}
