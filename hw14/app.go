package mwcRestAPI

import (
	"golang_beginning/hw14/controllers"
	"golang_beginning/hw14/models"
	"golang_beginning/hw14/view"
	"net/http"
	"github.com/gorilla/mux"
)

func Main() {
	models.LoadItems()
	router := mux.NewRouter()
	router.HandleFunc("/items", view.WithTemplate(controllers.GetItems)).Methods("GET")
	router.HandleFunc("/items", view.WithTemplate(controllers.CreateItem)).Methods("POST")
	router.HandleFunc("/items/{id}", view.WithTemplate(controllers.DeleteItem)).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
