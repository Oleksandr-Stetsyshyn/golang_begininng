package weather

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var cityWeather http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file1")
	}

	apiKey := os.Getenv("API_ACCU_WEATHER_KEY")
	if apiKey == "" {
		log.Fatal("API_ACCU_WEATHER_KEY is not set")
	}

	accuWeatherURL := os.Getenv("ACCU_WEATHER_URL")
	if accuWeatherURL == "" {
		log.Fatal("ACCU_WEATHER_URL is not set")
	}

	vars := mux.Vars(r)
	city := vars["city"]

	cityKey, err := getCityKey(apiKey, accuWeatherURL, city)
	if err != nil {
		log.Fatal(err)
	}

	currentWeather, err := getCurrentWeather(apiKey, accuWeatherURL, cityKey)
	if err != nil {
		log.Fatal(err)
	}

	response := fmt.Sprintf("{\"City\": \"%s\", \"Weather\": \"%s\", \"Temperature\": {\"Value\": %.0f, \"Unit\": \"%s\"}}", cases.Title(language.Und).String(city), currentWeather.WeatherText, currentWeather.Temperature.Metric.Value, currentWeather.Temperature.Metric.Unit)
	w.Write([]byte(response))
}
