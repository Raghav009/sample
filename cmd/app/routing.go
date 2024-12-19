package main

import (
	"database/sql"
	"net/http"
	"sample/internal/handlers"
	"sample/internal/middleware"

	"github.com/rs/cors"
)

func HandlerRouting(dbConn *sql.DB) http.Handler {
	// Define HTTP routes
	mux := http.NewServeMux()

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
	return handler
}
