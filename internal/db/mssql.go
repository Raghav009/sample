package db

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func MSSQLConnection(connString string) (*sql.DB, error) {
	if connString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// "Server=tcp:127.0.0.1,1433;Initial Catalog=AdventureWorks2016;Persist Security Info=False;User ID=ra;Password=root;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=True;"
