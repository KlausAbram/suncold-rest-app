package usecase

import (
	"context"
	"time"

	"github.com/klaus-abram/suncold-restful-app/api/external/cash"
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherCase struct {
	adapter *owmadapter.OwmAdapter
	store   storage.WeatherSearching
	cash    *cash.CashStorage
}

func NewWeatherCase(adapter *owmadapter.OwmAdapter, store storage.WeatherSearching, rdb *cash.CashStorage) *WeatherCase {
	return &WeatherCase{
		adapter: adapter,
		store:   store,
		cash:    rdb,
	}
}

func (cs *WeatherCase) GetWeatherCity(agentId int, location string) (*models.WeatherResponse, error) {
	dataResp, err := cs.adapter.GetOwmWeatherData(location)
	if err != nil {
		return nil, err
	}

	infId, agent, err := cs.store.PostWeatherData(agentId, *dataResp)
	if err != nil {
		return nil, err
	}

	err = cs.cash.SetRequestToCash(&models.WeatherRequest{
		Id:    infId,
		Date:  time.Now().Format("2006-01-02"),
		Mod:   "0",
		Agent: agent,
	}, context.Background())

	if err != nil {
		return nil, err
	}

	dataResp.InfId = infId

	return dataResp, nil
}
