package storage

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

var ErrorEmptyList = errors.New("empty list")

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
		return nil, ErrorEmptyList
	}

	return dataRequests, nil
}

func (wst *HistoryStorage) GetAgentHistoryData(agent string) ([]models.WeatherRequest, error) {
	var (
		dataAgent []models.WeatherRequest
		goalAgent string
	)

	tx, err := wst.db.Begin()
	if err != nil {
		return nil, err
	}

	queryAgent := fmt.Sprintf("SELECT (agent_name) FROM %s WHERE name=$1", agentsTable)

	if err := wst.db.Get(&goalAgent, queryAgent, agent); err != nil {
		if goalAgent == "" {
			return nil, errors.New("agent not found")
		}

		tx.Rollback()
		return nil, err
	}

	queryReq := fmt.Sprintf("SELECT * FROM %s WHERE author_name=$1", requestsTable)

	if err := wst.db.Select(&dataAgent, queryReq, agent); err != nil {
		tx.Rollback()
		return nil, err
	}

	if dataAgent == nil {
		return nil, ErrorEmptyList
	}

	return dataAgent, tx.Commit()
}
