package models

type Docente struct {
	ID              int    `json:"id"`
	Nombres         string `json:"nombres"`
	Apellidos       string `json:"apellidos"`
	Correo          string `json:"correo"`
	Telefono        string `json:"telefono"`
	Departamento    string `json:"departamento"`
	Especialidad    string `json:"especialidad"`
	TituloAcademico string `json:"titulo_academico"`
}