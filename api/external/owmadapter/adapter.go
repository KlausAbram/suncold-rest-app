package owmadapter

import "github.com/klaus-abram/suncold-restful-app/models"

var owmSet = struct{ Metric, Lang string }{Metric: "C", Lang: "RU"}

type OwmAdapter interface {
	GetOwmWeatherData(location string) (*models.WeatherParams, error)
}

type Adapter struct {
	OwmAdapter
}

func NewOwmAdapter() *Adapter {
	return &Adapter{
		OwmAdapter: NewWeatherAdapter(),
	}
}
