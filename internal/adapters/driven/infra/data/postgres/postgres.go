package postgres

import (
	"database/sql"
	"fmt"
	"maycms/internal/domain/entities"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
}

// Exec implements ports.Database.
func (p *PostgresDatabase) Exec(db *sql.DB, q string, c entities.Content) error {
	_, err := db.Exec(q, c.Title, c.ContentText, c.Status)
	return err
}

// Query implements ports.Database.
func (p *PostgresDatabase) Query(db *sql.DB, q string) (*sql.Rows, error) {
	rows, err := db.Query(q)
	return rows, err
}

// Query implements ports.Database.
func (p *PostgresDatabase) QueryRow(db *sql.DB, q string, id int) *sql.Row {
	row := db.QueryRow(q, id)
	return row
}

// CloseConnection implements ports.Database.
func (p *PostgresDatabase) CloseConnection(db *sql.DB) {
	db.Close()
}

// OpenConnection implements ports.Database.
func (p *PostgresDatabase) OpenConnection() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		panic("Failed recovering the environment variables")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()

	return db, err
}

func NewPostgresDB() *PostgresDatabase {
	return &PostgresDatabase{}
}
