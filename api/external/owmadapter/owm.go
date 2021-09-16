package owmadapter

import (
	"os"

	"github.com/briandowns/openweathermap"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type WeatherKeyStorage struct {
	apiWeatherKey string
}

func NewWeatherKeyStorage() *WeatherKeyStorage {
	return &WeatherKeyStorage{apiWeatherKey: os.Getenv("OWM_API_KEY")}
}

func (adp *WeatherKeyStorage) GetOwmWeatherData(location string) (*models.WeatherParams, error) {
	data, err := openweathermap.NewCurrent(owmSet.Lang, owmSet.Metric, adp.apiWeatherKey)
	if err != nil {
		return nil, err
	}

	var inputWeather = models.WeatherParams{
		Temperature: data.Main.Temp,
		Pressure:    data.Main.Pressure,
		Rain:        data.Rain.OneH,
		Cloud:       data.Clouds.All,
		WindSpeed:   data.Wind.Speed,
		Humidity:    data.Main.Humidity,
	}

	return &inputWeather, nil
}
