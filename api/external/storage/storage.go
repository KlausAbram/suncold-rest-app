package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

const (
	agentsTable   = "agents"
	requestsTable = "requests"
	statesTable   = "states"
	linksTable    = "links"
)

type Authorisation interface {
	CreateAgent(agent models.Agent) (int, error)
	GetAgent(agentname, password string) (int, error)
}

type WeatherSearching interface {
	PostWeatherData(agentId int, input models.WeatherParams) error
}

type GettingWeatherHistory interface {
}

type Storage struct {
	Authorisation
	WeatherSearching
	GettingWeatherHistory
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorisation:    NewAuthStorage(db),
		WeatherSearching: NewWeatherStorage(db),
	}
}
