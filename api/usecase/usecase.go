package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type Authorisation interface {
	CreateAgent(agent models.Agent) (int, error)
	CreateJWT(agentname, password string) (string, error)
	ParseJWT(token string) (int, error)
}

type WeatherSearching interface {
}

type GettingWeatherHistory interface {
}

type UseCase struct {
	Authorisation
	WeatherSearching
	GettingWeatherHistory
}

func NewUseCase(storage *storage.Storage) *UseCase {
	return &UseCase{
		Authorisation: NewAuthCase(storage.Authorisation),
	}
}
