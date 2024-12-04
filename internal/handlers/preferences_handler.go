package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sample/internal/db"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

func GetPreferencesHandler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	preferences, err := db.GetPreferences(dbConn)
	if err != nil {
		http.Error(w, "Error retrieving preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}
