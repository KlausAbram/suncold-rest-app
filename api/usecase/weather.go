package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherCase struct {
	adapter *owmadapter.Adapter
	store   storage.WeatherSearching
}

func NewWeatherCase(adapter *owmadapter.Adapter, store storage.WeatherSearching) *WeatherCase {
	return &WeatherCase{
		adapter: adapter,
		store:   store,
	}
}

func (cs *WeatherCase) GetWeatherCity(agentId int, location string) (*models.WeatherParams, error) {
	dataParams, err := cs.adapter.GetOwmWeatherData(location)
	if err != nil {
		return nil, err
	}

	/*if err := cs.store.PostWeatherData(*dataParams); err != nil {
		return nil, err
	} */
	return dataParams, nil
}
