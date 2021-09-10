package models

const (
	BaseSetting   = 0
	DetailSetting = 1
)

type ModWeatherSetting string
type Location string

type WeatherParams struct {
	Temperature int
	Pressure    int
	Rain        string
	Cloud       string
	Wind        string
}

type WeatherRequest struct {
	Id    int
	Date  string
	Mod   string ``
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
