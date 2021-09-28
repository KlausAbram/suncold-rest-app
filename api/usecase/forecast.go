package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type ForecastCase struct {
	adapter *owmadapter.OwmAdapter
}

func NewForecastCase(adp *owmadapter.OwmAdapter) *ForecastCase {
	return &ForecastCase{adapter: adp}
}

func (cs *ForecastCase) GettingForecastByDays(location string, days int) ([]models.WeatherResponse, error) {
	return nil, nil
}
