package weather

import (
	"encoding/json"
	"errors"
	"net/http"
)

type location struct {
	Key string `json:"Key"`
}

type currentWeather struct {
	WeatherText string `json:"WeatherText"`
	Temperature struct {
		Metric struct {
			Value float64 `json:"Value"`
			Unit  string  `json:"Unit"`
		} `json:"Metric"`
	} `json:"Temperature"`
}

func getCityKey(apiKey, accuWeatherURL, city string) (string, error) {
	getCityKeyURL := accuWeatherURL + "/locations/v1/cities/search?apikey=" + apiKey + "&q=" + city

	res, err := http.Get(getCityKeyURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var locations []location
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return "", err
	}

	if len(locations) == 0 {
		return "", errors.New("no locations found")
	}

	return locations[0].Key, nil
}

func getCurrentWeather(apiKey, accuWeatherURL, cityKey string) (*currentWeather, error) {
	url := accuWeatherURL + "/currentconditions/v1/" + cityKey + "?apikey=" + apiKey

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var current []currentWeather
	err = json.NewDecoder(res.Body).Decode(&current)
	if err != nil {
		return nil, err
	}

	if len(current) == 0 {
		return nil, errors.New("no current weather found")
	}

	return &current[0], nil
}
