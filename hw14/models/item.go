package models

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Category    string  `json:"category"`
}

var Items []Item

func LoadItems() {
	file, err := os.ReadFile("hw14/db/items.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &Items)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveItems() {
	file, err := json.MarshalIndent(Items, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./hw14/db/items.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
