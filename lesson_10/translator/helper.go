package translator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type TranslationResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Translated  string `json:"translated"`
		Text        string `json:"text"`
		Translation string `json:"translation"`
	} `json:"contents"`
}

var translationHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	text := string(body)

	if text == "" {
		http.Error(w, "Empty text", http.StatusBadRequest)
		return
	}

	translation, err := getTranslate(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"Translation": translation}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getTranslate(text string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file1")
	}

	apiKey := os.Getenv("FUN_TRANSLATIONS_URL")
	if apiKey == "" {
		log.Fatal("FUN_TRANSLATIONS_URL is not set")
	}

	encodedText := url.QueryEscape(text)

	url := apiKey + "/mandalorian?text=" + encodedText

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var response TranslationResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Contents.Translated, nil
}
