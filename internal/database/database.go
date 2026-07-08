package database

import (
	"fmt"
	"os"
	"time"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	fmt.Println("HOST:", host)
	fmt.Println("PORT:", port)
	fmt.Println("USER:", user)
	fmt.Println("DB:", dbname)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)
    
    var db *gorm.DB
    var err error
    
    for i := 0; i < 10; i++ {
        
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        
        if err == nil {
            
            sqlDB, err2 := db.DB()
            
            if err2 == nil && sqlDB.Ping() == nil {
            
                fmt.Println("Conectado a PostgreSQL")
                break
            }
        }

        fmt.Println("Esperando a PostgreSQL...")

        time.Sleep(3 * time.Second)
    }
    
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Docente{},
		&models.SesionTutoria{},
		&models.Asistencia{},
		&models.Evidencia{},
		&models.TipoTutoria{},
		&models.SolicitudTutoria{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}