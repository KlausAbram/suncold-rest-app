package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherStorage struct {
	db *sqlx.DB
}

func NewWeatherStorage(db *sqlx.DB) *WeatherStorage {
	return &WeatherStorage{db: db}
}

func (wst *WeatherStorage) PostWeatherData(agentId int, input models.WeatherResponse) (int, error) {
	tx, err := wst.db.Begin()
	if err != nil {
		return 0, err
	}

	var (
		reqId     int
		agentName string
	)

	queryAgent := fmt.Sprintf("SELECT (agent_name) FROM %s WHERE id=$1", agentsTable)
	err = wst.db.Get(&agentName, queryAgent, agentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	queryRequests := fmt.Sprintf("INSERT INTO %s (author_name) VALUES ($1) RETURNING id", requestsTable)
	rowName := tx.QueryRow(queryRequests, agentName)
	if err := rowName.Scan(&reqId); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryStates := fmt.Sprintf("INSERT INTO %s (location, temperature, pressure, rain, clouds, wind) VALUES ($1, $2, $3, $4, $5, $6)",
		statesTable)
	_, err = tx.Exec(queryStates, input.Location, input.Temperature, input.Pressure, input.Rain, input.Cloud, input.WindSpeed)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	queryLinks := fmt.Sprintf("INSERT INTO %s (request_id, state_id, agent_id) VALUES ($1, $2, $3)", linksTable)
	_, err = tx.Exec(queryLinks, reqId, reqId, agentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return reqId, tx.Commit()

}
