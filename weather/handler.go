package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HandleWeatherRequest(w http.ResponseWriter, r *http.Request, svc WeatherService) {
	path := strings.TrimPrefix(r.URL.Path, "/weather/")
	log.Printf("Received request for path: %q", r.URL.Path)
	coords, err := parseCoordinates(path)
	if err != nil {
		http.Error(w, "Invalid coordinates format. Expected /weather/{lat},{lon}", http.StatusBadRequest)
		return
	}

	brief, err := svc.Forecast(coords)
	if err != nil {
		message := fmt.Sprintf("Forecast unavailable because: %v", err)
		http.Error(w, message, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brief)
}

func parseCoordinates(path string) (Coordinates, error) {
	parts := strings.Split(path, ",")
	if len(parts) != 2 {
		return Coordinates{}, errors.New("invalid coordinate format")
	}

	lat, err1 := strconv.ParseFloat(parts[0], 32)
	lon, err2 := strconv.ParseFloat(parts[1], 32)

	if err1 != nil || err2 != nil {
		return Coordinates{}, fmt.Errorf("could not parse lat/lon: %w, %w", err1, err2)
	}

	return Coordinates{
		Latitude:  float32(lat),
		Longitude: float32(lon),
	}, nil
}
