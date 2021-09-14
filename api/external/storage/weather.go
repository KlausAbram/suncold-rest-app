package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherStorage struct {
	db *sqlx.DB
}

func NewWeatherStorage(db *sqlx.DB) *WeatherStorage {
	return &WeatherStorage{db: db}
}

func (wst *WeatherStorage) PostWeatherData(agentId int, input models.WeatherParams) error {
	return nil
}
