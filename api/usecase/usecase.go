package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type Authorisation interface {
	CreateAgent(agent models.Agent) (int, error)
	CreateJWT(agentname, password string) (string, error)
	ParseJWT(token string) (int, error)
}

type WeatherSearching interface {
	GetWeatherCity(agentId int, location string) (*models.WeatherParams, error)
}

type GettingWeatherHistory interface {
}

type UseCase struct {
	Authorisation
	WeatherSearching
	GettingWeatherHistory
}

func NewUseCase(adapter *owmadapter.OwmAdapter, store *storage.Storage) *UseCase {
	return &UseCase{
		Authorisation:    NewAuthCase(&store.Authorisation),
		WeatherSearching: NewWeatherCase(adapter, store.WeatherSearching),
	}
}
