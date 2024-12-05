package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sample/internal/auth"
	"sample/internal/db"
	"sample/internal/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode the login request body
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real-world app, you should validate the credentials against a database.
	// Here, we hard-code the username and password check for simplicity.

	user, err := db.GetUser(credentials.Username, dbConn)
	if err != nil {
		http.Error(w, "Error retrieving User", http.StatusInternalServerError)
		return
	}
	if user == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not available"))
		return
	}
	err = auth.ComparePassword(credentials.Password, user.Password)
	if err == nil {
		// Generate JWT token for the user
		token, err := auth.GenerateJWT(credentials.Username, "user")
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Send the token back as a response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
