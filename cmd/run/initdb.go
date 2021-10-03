package run

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// InitPostgresStorage make in main
func InitPostgresStorage() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("CONNECT_STRING"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
