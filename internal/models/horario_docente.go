package models

type HorarioDocente struct {
	ID         int    `json:"id"`
	DocenteID  int    `json:"docente_id"`
	MateriaID  int    `json:"materia_id"`
	DiaSemana  string `json:"dia_semana"`
	HoraInicio string `json:"hora_inicio"`
	HoraFin    string `json:"hora_fin"`
	Modalidad  string `json:"modalidad"`
	Aula       string `json:"aula"`
}
