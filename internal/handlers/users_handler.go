package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sample/internal/auth"
	"sample/internal/db"
	"sample/internal/models"
	"sample/internal/utils"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := db.GetUsers(dbConn)
	if err != nil {
		http.Error(w, "Error retrieving users"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"IsAdmin"`
	}

	// Decode the login request body
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(credentials.Username, dbConn)
	if err != nil {
		http.Error(w, "Error retrieving User", http.StatusInternalServerError)
		return
	}
	if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}
	hash, err := auth.CreateHashedPassword(credentials.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Error hashing plain password"))
		return
	}
	_user := models.User{
		UserName: credentials.Username,
		Password: hash,
		IsAdmin:  credentials.IsAdmin,
	}

	err = db.AddUser(_user, dbConn)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Error Adding User"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User Added Successfully"}`))
}
