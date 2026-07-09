package models

type Usuario struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Nombre       string `gorm:"not null" json:"nombre"`
	Correo       string `gorm:"uniqueIndex;not null" json:"correo"`
	PasswordHash string `gorm:"not null" json:"-"`
	Rol          string `gorm:"not null" json:"rol"`
}
