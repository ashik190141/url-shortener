package db

import (
	"log"
	"os"
	"url-shortener/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	// Build DSN from docker-compose environment variables
	if dsn == "" {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		
		if host != "" {
			dsn = "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
		} else {
			log.Fatalln("Error: DATABASE_URL or POSTGRES_* environment variables are not set.")
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database using GORM: %v", err)
	}

	log.Println("Successfully connected to the database with GORM!")

	err = db.AutoMigrate(&model.URL{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate the database schema: %v", err)
	}
	log.Println("Database migration completed successfully!")
	
	return db
}
