package models

const (
	BaseSetting   = 0
	DetailSetting = 1
)

type ModWeatherSetting string
type Location string

type WeatherResponse struct {
	Temperature int    `json:"temperature" db:"temperature"`
	Pressure    int    `json:"pressure" db:"pressure"`
	Rain        int    `json:"rain" db:"rain"`
	Cloud       int    `json:"cloud" db:"clouds"`
	WindSpeed   int    `json:"wind_speed" db:"wind"`
	Humidity    int    `json:"humidity"`
	Location    string `json:"location" db:"location"`
	InfId       int    `json:"inf_id" db:"id"`
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
