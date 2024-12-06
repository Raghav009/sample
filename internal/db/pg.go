package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PGConnection(connString string) (*sql.DB, error) {
	if connString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db.DB, nil
}

// "host=localhost port=5432 user=postgres password=password dbname=Sample sslmode=disable"
