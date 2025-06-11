package weather

import "time"

type TempDescr string
type ConditionsDescr string
type PeriodName string

type BriefWeather struct {
	Period      PeriodName      `json:"period"`
	Temperature TempDescr       `json:"temperature"`
	Conditions  ConditionsDescr `json:"conditions"`
}

type Coordinates struct {
	Latitude  string
	Longitude string
}

type WeatherService interface {
	Forecast(coords Coordinates) (BriefWeather, error)
}

// returned from the final forecast URL
type Period struct {
	Name          PeriodName      `json:"name"` // e.g. "Today", "Tonight"
	StartTime     time.Time       `json:"startTime"`
	EndTime       time.Time       `json:"endTime"`
	Temperature   int             `json:"temperature"`
	ShortForecast ConditionsDescr `json:"shortForecast"`
}
