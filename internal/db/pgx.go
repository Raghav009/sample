package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PGXConnection(connString string) (*pgxpool.Pool, error) {
	if connString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	// Set up a connection pool
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Optional: Adjust connection settings if needed
	config.MaxConns = 10                               // Example: maximum number of connections
	config.ConnConfig.ConnectTimeout = 5 * time.Second // Timeout for connection

	// Create a connection pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
