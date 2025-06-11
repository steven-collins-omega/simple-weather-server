package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// template for NWS points API, which we call to get the "real" forecast URL
	urlTemplate = "https://api.weather.gov/points/%s,%s"

	// in degrees Fahrenheit
	minHot  = 75
	maxCold = 50
)

func getForecastURL(ctx context.Context, coords Coordinates) (string, error) {
	url := fmt.Sprintf(urlTemplate, coords.Latitude, coords.Longitude)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Properties struct {
			Forecast string `json:"forecast"`
		} `json:"properties"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Properties.Forecast, nil
}

func fetchForecast(ctx context.Context, forecastURL string) ([]Period, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, forecastURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Properties struct {
			Periods []Period `json:"periods"`
		} `json:"properties"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Properties.Periods, nil
}

func extractBriefForecast(periods []Period) (bw BriefWeather, ok bool) {
	for _, p := range periods {
		if p.Name == "Today" {
			return BriefWeather{
				Period:      PeriodName(p.Name),
				Temperature: mapTemperature(p.Temperature),
				Conditions:  ConditionsDescr(p.ShortForecast),
			}, true
		}
	}

	// sometimes there is no "Today" in which case we explicitly check the times
	now := time.Now()
	for _, p := range periods {
		if now.After(p.StartTime) && now.Before(p.EndTime) {
			return BriefWeather{
				Period:      PeriodName(p.Name),
				Temperature: mapTemperature(p.Temperature),
				Conditions:  ConditionsDescr(p.ShortForecast),
			}, true
		}
	}

	return BriefWeather{}, false
}

func mapTemperature(temp int) TempDescr {
	switch {
	case temp < maxCold:
		return "cold"
	case temp > minHot:
		return "hot"
	default:
		return "moderate"
	}
}
