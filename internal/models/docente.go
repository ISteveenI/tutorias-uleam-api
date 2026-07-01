package models

type Docente struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Nombres         string `gorm:"not null" json:"nombres"`
	Apellidos       string `gorm:"not null" json:"apellidos"`
	Correo          string `gorm:"unique;not null" json:"correo"`
	Telefono        string `json:"telefono"`
	Departamento    string `json:"departamento"`
	Especialidad    string `json:"especialidad"`
	TituloAcademico string `json:"titulo_academico"`
}
