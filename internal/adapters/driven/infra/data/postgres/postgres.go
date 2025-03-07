package postgres

import (
	"database/sql"
	"fmt"
	"maycms/internal/domain/entities"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
}

// Exec implements data.Database.
func (p *PostgresDatabase) Exec(db *sql.DB, q string, c entities.Content) error {
	_, err := db.Exec(q, c.Title, c.ContentText, c.Status)
	return err
}

// Query implements data.Database.
func (p *PostgresDatabase) Query(db *sql.DB, q string) (*sql.Rows, error) {
	rows, err := db.Query(q)
	return rows, err
}

// Query implements data.Database.
func (p *PostgresDatabase) QueryRow(db *sql.DB, q string, id int) *sql.Row {
	row := db.QueryRow(q, id)
	return row
}

// CloseConnection implements data.Database.
func (p *PostgresDatabase) CloseConnection(db *sql.DB) {
	db.Close()
}

// OpenConnection implements data.Database.
func (p *PostgresDatabase) OpenConnection() (*sql.DB, error) {
	dsn := "postgres://gon:a12345z@localhost:5432/maycms-db?sslmode=disable"
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
