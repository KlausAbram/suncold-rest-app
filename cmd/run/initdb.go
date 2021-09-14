package run

import (
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
)

// InitPostgresStorage make in main
func InitPostgresStorage() (*sqlx.DB, error) {

	db, errDB := storage.NewPostgresStorage()
	if errDB != nil {
		return nil, errDB
	}

	return db, nil
}
