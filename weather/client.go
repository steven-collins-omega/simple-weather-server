package weather

import "fmt"

type NationalWeatherService struct{}

func (n NationalWeatherService) Forecast(coords Coordinates) (BriefWeather, error) {
	url, err := getForecastURL(coords)
	if err != nil {
		return BriefWeather{}, fmt.Errorf("error accessing National Weather Service: %w", err)
	}

	periods, err := fetchForecast(url)
	if err != nil {
		return BriefWeather{}, fmt.Errorf("no forecasts available for location. Error: %w", err)
	}

	brief, ok := extractBriefForecast(periods);
	if !ok {
		return BriefWeather{}, fmt.Errorf("today's forecast not available")
	}

	return brief, nil
}
