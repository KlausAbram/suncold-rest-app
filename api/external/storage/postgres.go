package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/cmd/init"
)

func NewPostgresStorage(cfg init.DataBaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
