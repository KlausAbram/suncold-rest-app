package storage

import (
	"github.com/jmoiron/sqlx"
	"os"
)

func NewPostgresStorage() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("CONNECT_STRING"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
