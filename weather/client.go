package weather

import "fmt"

type NationalWeatherService struct{}

func (n NationalWeatherService) Forecast(coords Coordinates) (BriefWeather, error) {
	url, err := getForecastURL(coords)
	if err != nil {
		return BriefWeather{}, fmt.Errorf("error accessing National Weather Service: %w", err)
	}

	periods, err := fetchForecast(url)
	if err != nil || len(periods) == 0 {
		return BriefWeather{}, fmt.Errorf("no forecasts available for location. error: %w", err)
	}

	brief, ok := extractBriefForecast(periods);
	if !ok {
		return BriefWeather{}, fmt.Errorf("no forecast for 'Today' or current time period")
	}

	return brief, nil
}
