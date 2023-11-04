package mwcRestAPI

import (
	"golang_beginning/hw14/controllers"
	"golang_beginning/hw14/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	models.LoadItems()
	router := mux.NewRouter()
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
