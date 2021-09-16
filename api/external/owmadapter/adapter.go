package owmadapter

import "github.com/klaus-abram/suncold-restful-app/models"

var owmSet = struct{ Metric, Lang string }{Metric: "C", Lang: "RU"}

type OwmInterface interface {
	GetOwmWeatherData(location string) (*models.WeatherParams, error)
}

type OwmAdapter struct {
	OwmInterface
}

func NewOwmAdapter() *OwmAdapter {
	return &OwmAdapter{
		OwmInterface: NewWeatherKeyStorage(),
	}
}
