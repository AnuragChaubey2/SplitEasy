package database

import (
	"log"
	"os"
	"spliteasy/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("AutoMigration failed: %v", err)
	}

	DB = db
	log.Println("Database connection established successfully.")
}
