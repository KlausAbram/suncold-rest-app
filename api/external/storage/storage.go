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
	PostWeatherData(agentId int, input models.WeatherResponse) (int, error)
}

type GettingWeatherHistory interface {
	GetHistoryLocationData(location string) ([]models.WeatherResponse, error)
	GetHistoryMomentData(moment string) ([]models.WeatherRequest, error)
	GetAgentHistoryData(agent string) ([]models.WeatherRequest, error)
}

type Storage struct {
	Authorisation
	WeatherSearching
	GettingWeatherHistory
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorisation:         NewAuthStorage(db),
		WeatherSearching:      NewWeatherStorage(db),
		GettingWeatherHistory: NewHistoryStorage(db),
	}
}
