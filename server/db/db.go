package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/go-chat?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() {
	db.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
