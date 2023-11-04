// controllers/itemController.go
package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "golang_beginning/hw14/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Items)
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
    json.NewEncoder(w).Encode(newItem)
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
    json.NewEncoder(w).Encode(models.Items)
}