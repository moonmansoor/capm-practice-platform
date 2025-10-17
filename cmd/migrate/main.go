package main

import (
	"context"
	"log"

	"capm-exam-system/internal/database"
)

func main() {
	// Connect to database
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create tables
	ctx := context.Background()
	if err := db.CreateTables(ctx); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	log.Println("Database migration completed successfully")
}
