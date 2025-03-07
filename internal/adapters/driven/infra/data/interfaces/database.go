package data

import (
	"database/sql"
	"maycms/internal/domain/entities"
)

type Database interface {
	OpenConnection() (*sql.DB, error)
	CloseConnection(db *sql.DB)
	QueryRow(db *sql.DB, q string, id int) *sql.Row
	Query(db *sql.DB, q string) (*sql.Rows, error)
	Exec(db *sql.DB, q string, c entities.Content) error
}
