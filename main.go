package main

import (
	"log"
	"net/http"

	"github.com/steven-collins-omega/simple-weather-server/weather"
)

func main() {
	svc := weather.NationalWeatherService{}

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		weather.HandleWeatherRequest(w, r, svc)
	})

	log.Println("Server running on :8080. Try: http://localhost:8080/weather/40.6782,-73.9442")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
