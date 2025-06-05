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
	Forecast(coords Coordinates) BriefWeather
}
