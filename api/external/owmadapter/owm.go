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

func (adp *WeatherKeyStorage) GetOwmWeatherData(location string) (*models.WeatherResponse, error) {
	client, err := openweathermap.NewCurrent(owmSet.Metric, owmSet.Lang, adp.apiWeatherKey)
	if err != nil {
		return nil, err
	}

	if err := client.CurrentByName(location); err != nil {
		return nil, err
	}

	var inputWeather = &models.WeatherResponse{
		Temperature: int(client.Main.Temp),
		Pressure:    int(client.Main.Pressure),
		Rain:        int(client.Rain.OneH),
		Cloud:       int(client.Clouds.All),
		WindSpeed:   int(client.Wind.Speed),
		Humidity:    int(client.Main.Humidity),
		Location:    location,
	}

	//return &inputWeather, nil
	return inputWeather, nil
}

func (adp *WeatherKeyStorage) GetForecastInfo(location string) ([]models.WeatherResponse, error) {
	return nil, nil
}
