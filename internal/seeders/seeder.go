package seeders

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
)

func SeedUsers(db *gorm.DB) {

	ctx := context.Background()

	users := []models.User{
		{
			Username: "admin",
			Password: "admin123",
			Role:     "admin",
		},
		{
			Username: "docente",
			Password: "docente123",
			Role:     "docente",
		},
		{
			Username: "estudiante",
			Password: "estudiante123",
			Role:     "estudiante",
		},
	}

	for _, u := range users {

		var count int64

		db.WithContext(ctx).
			Model(&models.User{}).
			Where("username = ?", u.Username).
			Count(&count)

		if count > 0 {
			continue
		}

		hash, _ := bcrypt.GenerateFromPassword(
			[]byte(u.Password),
			bcrypt.DefaultCost,
		)

		u.Password = string(hash)

		db.Create(&u)

		fmt.Println("Usuario creado:", u.Username)
	}
}
