package controllers

import (
	"encoding/json"
	"net/http"
	"golang_beginning/hw14/models"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.Items = append(models.Items, newItem)
	models.SaveItems()

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range models.Items {
		if item.ID == params["id"] {
			models.Items = append(models.Items[:index], models.Items[index+1:]...)
			models.SaveItems()
			break
		}
	}

}
