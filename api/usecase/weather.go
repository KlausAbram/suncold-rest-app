package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherCase struct {
	adapter *owmadapter.OwmAdapter
	store   storage.WeatherSearching
}

func NewWeatherCase(adapter *owmadapter.OwmAdapter, store storage.WeatherSearching) *WeatherCase {
	return &WeatherCase{
		adapter: adapter,
		store:   store,
	}
}

func (cs *WeatherCase) GetWeatherCity(agentId int, location string) (*models.WeatherResponse, error) {
	dataResp, err := cs.adapter.GetOwmWeatherData(location)
	if err != nil {
		return nil, err
	}

	infId, err := cs.store.PostWeatherData(agentId, *dataResp)
	if err != nil {
		return nil, err
	}

	dataResp.InfId = infId

	return dataResp, nil
}
