package main

import (
	"fmt"
	"log"
	"movie-reccomendation-api/config"
	"movie-reccomendation-api/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := database.NewDB(cfg.DBConnectionString())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Test database connection
	err = db.Exec("SELECT 1").Error
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("Database connection successful!")
}
