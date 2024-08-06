package db

import (
	"github.com/Mogza/FandC/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Init() *gorm.DB {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Artist{}, &models.Music{}, &models.Token{}, &models.Transaction{}, &models.Royalty{})
	if err != nil {
		log.Fatalf("Failed to migrate database : %v", err)
	}

	return db
}
