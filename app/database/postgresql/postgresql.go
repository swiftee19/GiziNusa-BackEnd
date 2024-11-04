package postgresql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Construct the connection string with spaces between parameters
	destination := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRESQLHOST"),
		os.Getenv("POSTGRESQLUSER"),
		os.Getenv("POSTGRESQLPASSWORD"),
		os.Getenv("POSTGRESQLDATABASENAME"),
		os.Getenv("POSTGRESQLDATABASEPORT"),
	)

	// Open the connection
	db, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Database connected successfully!")
	DB = db // Assign to global DB variable

	return DB, nil
}