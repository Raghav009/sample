package db

import (
	"database/sql"
	"sample/internal/models"
)

func GetPreferences(db *sql.DB) ([]models.Preferences, error) {
	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var preferences []models.Preferences
	for rows.Next() {
		var pref models.Preferences
		if err := rows.Scan(&pref.UserId, &pref.PageName, &pref.ViewPreferences); err != nil {
			return nil, err
		}
		preferences = append(preferences, pref)
	}
	return preferences, nil
}
