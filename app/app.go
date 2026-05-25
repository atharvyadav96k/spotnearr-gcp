package app

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is missing")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database via GORM: %v", err)
	}

	log.Println("Database connection successfully established.")
	err = DB.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed to run AutoMigration: %v", err)
	}

	log.Println("Database migration completed successfully.")
}
