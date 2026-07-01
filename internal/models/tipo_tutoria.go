package models

type TipoTutoria struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
