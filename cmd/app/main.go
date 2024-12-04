package main

import (
	"fmt"
	"log"
	"net/http"

	"sample/internal/config"
	"sample/internal/db"
	"sample/internal/handlers"
	"sample/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	dbConn, err := db.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("Connected to the database!")

	// Define HTTP routes
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/preferences", middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPreferencesHandler(w, r, dbConn)
	}))
	http.HandleFunc("/login", handlers.LoginHandler)

	// Start the HTTP server
	log.Printf("Starting server on %s...\n", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// del app.exe; go build ./cmd/app; ./app.exe
