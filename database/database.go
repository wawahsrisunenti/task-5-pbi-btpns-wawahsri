// database/database.go

package database

import (
	"fmt"
	"log"
	"os"

	"PBI_BTPN/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB menginisialisasi koneksi ke PostgreSQL
func InitDB() *gorm.DB {
	// Load variabel lingkungan dari file .env
	loadEnv()

	// Dapatkan konfigurasi koneksi dari variabel lingkungan
	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Gorm Logger mode (Silent, Error, Warn, Info)
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Photo{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return db
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
