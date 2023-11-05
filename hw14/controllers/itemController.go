package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang_beginning/hw14/models"
	"html/template"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	withTemplate(func(w http.ResponseWriter, r *http.Request) {

	}, "hw14/view/items.html")(w, r)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	withTemplate(func(w http.ResponseWriter, r *http.Request) {
		var newItem models.Item
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.Items = append(models.Items, newItem)
		models.SaveItems()

	}, "hw14/view/items.html")(w, r)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	withTemplate(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for index, item := range models.Items {
			if item.ID == params["id"] {
				models.Items = append(models.Items[:index], models.Items[index+1:]...)
				models.SaveItems()
				break
			}
		}

	}, "hw14/view/items.html")(w, r)
}

func withTemplate(handler func(http.ResponseWriter, *http.Request), templatePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handler(w, r)
		err = t.Execute(w, models.Items)
		if err != nil {
			return
		}
	}
}
