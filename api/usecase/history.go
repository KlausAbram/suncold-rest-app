package usecase

import (
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type HistoryCase struct {
	store storage.GettingWeatherHistory
}

func NewHistoryCase(store storage.GettingWeatherHistory) *HistoryCase {
	return &HistoryCase{store: store}
}

func (cs *HistoryCase) GetHistoryLocation(location string) ([]models.WeatherResponse, error) {
	dataCaseResponse, err := cs.store.GetHistoryLocationData(location)
	if err != nil {
		return nil, err
	}

	return dataCaseResponse, nil
}

func (cs *HistoryCase) GetHistoryMoment(moment string) ([]models.WeatherRequest, error) {
	dataCaseResponse, err := cs.store.GetHistoryMomentData(moment)
	if err != nil {
		return nil, err
	}

	return dataCaseResponse, nil
}
