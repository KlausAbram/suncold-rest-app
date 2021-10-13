package usecase

import (
	"context"

	"github.com/klaus-abram/suncold-restful-app/api/external/cash"
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type Authorisation interface {
	CreateAgent(agent models.Agent) (int, error)
	CreateJWT(agentname, password string) (string, error)
	ParseJWT(token string) (int, error)
}

type WeatherSearching interface {
	GetWeatherCity(agentId int, location string) (*models.WeatherResponse, error)
}

type GettingWeatherHistory interface {
	GetHistoryLocation(location string) ([]models.WeatherResponse, error)
	GetHistoryMoment(moment string) ([]models.WeatherRequest, error)
	GetAgentHistory(agent string) ([]models.WeatherRequest, error)
}

type GettingForecastByDays interface {
	GetForcastByDays(location string, days int) ([]models.WeatherResponse, error)
}

type GettingCashedData interface {
	GetCashedRequests(ctx context.Context) (*[]models.WeatherRequest, error)
}

type UseCase struct {
	Authorisation
	WeatherSearching
	GettingWeatherHistory
	GettingCashedData
}

func NewUseCase(adapter *owmadapter.OwmAdapter, store *storage.Storage, rdb *cash.CashStorage) *UseCase {
	return &UseCase{
		Authorisation:         NewAuthCase(&store.Authorisation),
		WeatherSearching:      NewWeatherCase(adapter, store.WeatherSearching, rdb),
		GettingWeatherHistory: NewHistoryCase(store.GettingWeatherHistory),
		GettingCashedData:     NewCashStorage(rdb),
	}
}
