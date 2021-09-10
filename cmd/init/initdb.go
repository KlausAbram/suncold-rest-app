package init

import (
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/api/storage"
	"github.com/spf13/viper"
	"os"
)

// InitPostgresStorage make in main
func InitPostgresStorage() (*sqlx.DB, error) {

	db, errDB := storage.NewPostgresStorage(DataBaseConfig{
		Name:     viper.GetString("db.name"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if errDB != nil {
		return nil, errDB
	}

	return db, nil
}
