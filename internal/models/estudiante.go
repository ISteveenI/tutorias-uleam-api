package models

type Estudiante struct {
	ID         int    `json:"id"`
	Nombres    string `json:"nombres"`
	Apellidos  string `json:"apellidos"`
	Correo     string `json:"correo"`
	Telefono   string `json:"telefono"`
	Carrera    string `json:"carrera"`
	Semestre   int    `json:"semestre"`
	Matricula  string `json:"matricula"`
}