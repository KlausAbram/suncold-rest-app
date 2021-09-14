package models

const (
	BaseSetting   = 0
	DetailSetting = 1
)

type ModWeatherSetting string
type Location string

type WeatherParams struct {
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
	Rain        float64 `json:"rain"`
	Cloud       int     `json:"cloud"`
	WindSpeed   float64 `json:"wind_speed"`
	Humidity    int     `json:"humidity"`
}

type WeatherRequest struct {
	Id    int
	Date  string
	Mod   string
	Agent string
}

type DataWeatherState struct {
	Id       int
	Location Location
	Mod      ModWeatherSetting
}

type RequestState struct {
	RequestId int
	StateId   int
	AgentId   int
}
