package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// template for NWS points API, which we call to get the "real" forecast URL
	urlTemplate = "https://api.weather.gov/points/%.4f,%.4f"

	// in degrees Fahrenheit
	minHot   = 75
	maxCold  = 50
)

func getForecastURL(coords Coordinates) (string, error) {
	url := fmt.Sprintf(urlTemplate, coords.Latitude, coords.Longitude)
	resp, err := http.Get(url)
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

func fetchForecast(forecastURL string) ([]Period, error) {
	resp, err := http.Get(forecastURL)
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
				Temperature: mapTemperature(p.Temperature),
				Conditions:  ConditionsDescr(p.ShortForecast),
			}, true
		}
	}

	// fail if "Today" is not present
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
