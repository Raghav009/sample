package main

import (
	"fmt"
	"log"
	"net/http"

	"sample/internal/config"
	"sample/internal/db"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	dbConn, err := db.PGConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("Connected to the database!")

	handler := HandlerRouting(dbConn)
	// Start the HTTP server
	log.Printf("Starting server on %s...\n", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// del app.exe; go build ./cmd/app; ./app.exe
