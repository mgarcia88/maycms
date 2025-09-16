package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDB() (*PostgresDatabase, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("missing DSN environment variable")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDatabase{DB: db}, nil
}

func (p *PostgresDatabase) Close() error {
	return p.DB.Close()
}

func (p *PostgresDatabase) GetDB() *sql.DB {
	return p.DB
}
