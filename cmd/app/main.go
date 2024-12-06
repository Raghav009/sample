package main

import (
	"fmt"
	"log"
	"net/http"

	"sample/internal/config"
	"sample/internal/db"
	"sample/internal/handlers"
	"sample/internal/middleware"

	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	// Define HTTP routes
	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/users", middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsersHandler(w, r, dbConn)
	}))
	mux.HandleFunc("/login", (func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, dbConn)
	}))
	mux.HandleFunc("/register", (func(w http.ResponseWriter, r *http.Request) {
		handlers.AddUserHandler(w, r, dbConn)
	}))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200/",
			"*",
		},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "*"},
		AllowCredentials: true,
	})
	handler := cors.Handler(mux)
	// Start the HTTP server
	log.Printf("Starting server on %s...\n", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// del app.exe; go build ./cmd/app; ./app.exe
