package storage

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type HistoryStorage struct {
	db *sqlx.DB
}

func NewHistoryStorage(db *sqlx.DB) *HistoryStorage {
	return &HistoryStorage{db: db}
}

func (wst *HistoryStorage) GetHistoryLocationData(location string) ([]models.WeatherResponse, error) {
	var dataStates []models.WeatherResponse

	query := fmt.Sprintf("SELECT * FROM %s WHERE location=$1", statesTable)

	if err := wst.db.Select(&dataStates, query, location); err != nil {
		return nil, err
	}

	return dataStates, nil
}

func (wst *HistoryStorage) GetHistoryMomentData(moment string) ([]models.WeatherRequest, error) {
	var dataRequests []models.WeatherRequest

	query := fmt.Sprintf("SELECT * FROM %s WHERE date=$1", requestsTable)

	if err := wst.db.Select(&dataRequests, query, moment); err != nil {
		return nil, err
	}

	if dataRequests == nil {
		return nil, errors.New("list is empty")
	}

	return dataRequests, nil
}