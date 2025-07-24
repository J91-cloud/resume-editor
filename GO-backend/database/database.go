package database

import (
	"database/sql"
)

type database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *database {
	return &database{
		db: db,
	}
}

func (d *database) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS jobs(
	
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	date_applied TEXT NOT NULL,
	job_type TEXT NOT NULL
	
	
	);
	`
	_, err := d.db.Exec(query)
	return err

}
