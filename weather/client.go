package weather

import (
	"context"
	"fmt"
)

type NationalWeatherService struct{}

func (n NationalWeatherService) Forecast(ctx context.Context, coords Coordinates) (BriefWeather, error) {
	url, err := getForecastURL(ctx, coords)
	if err != nil {
		return BriefWeather{}, fmt.Errorf("error accessing National Weather Service: %w", err)
	}

	periods, err := fetchForecast(ctx, url)
	if err != nil || len(periods) == 0 {
		return BriefWeather{}, fmt.Errorf("no forecasts available for location. error: %w", err)
	}

	brief, ok := extractBriefForecast(periods)
	if !ok {
		return BriefWeather{}, fmt.Errorf("no forecast for 'Today' or current time period")
	}

	return brief, nil
}
