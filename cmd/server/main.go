package main

import (
	"log"
	"net/http"
	"os"

	"capm-exam-system/internal/database"
	"capm-exam-system/internal/handlers"
	"capm-exam-system/internal/pdf"
	"capm-exam-system/internal/repository"
	"capm-exam-system/internal/service"
)

func main() {
	// Connect to database
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize layers
	repo := repository.New(db)
	svc := service.New(repo)
	pdfSvc := pdf.New()
	handlers := handlers.New(svc, pdfSvc)

	// Setup routes
	router := handlers.SetupRoutes()

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
