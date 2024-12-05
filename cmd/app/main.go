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
	http.Handle("/", middleware.CORS(http.DefaultServeMux))
	// Define HTTP routes
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/users", middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsersHandler(w, r, dbConn)
	}))
	http.HandleFunc("/login", (func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, dbConn)
	}))
	http.HandleFunc("/register", (func(w http.ResponseWriter, r *http.Request) {
		handlers.AddUserHandler(w, r, dbConn)
	}))

	// Start the HTTP server
	log.Printf("Starting server on %s...\n", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// del app.exe; go build ./cmd/app; ./app.exe
