package data

import "database/sql"

type Database interface {
	OpenConnection() (*sql.DB, error)
	CloseConnection(db *sql.DB)
	QueryRow(db *sql.DB, q string, id int) *sql.Row
}
