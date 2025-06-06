package weather

type TempDescr string
type ConditionsDescr string

type BriefWeather struct {
	Temperature TempDescr
	Conditions  ConditionsDescr
}

type Coordinates struct {
	Latitude  float32
	Longitude float32
}

type WeatherService interface {
	Forecast(coords Coordinates) (BriefWeather, error)
}

// returned from the final forecast URL
type Period struct {
	Name          string `json:"name"` // e.g. "Today", "Tonight"
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}