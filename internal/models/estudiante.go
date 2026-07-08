package models

type Estudiante struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Correo    string `json:"correo"`
	Telefono  string `json:"telefono"`
	Carrera   string `json:"carrera"`
	Semestre  int    `json:"semestre"`
	Matricula string `json:"matricula"`
}