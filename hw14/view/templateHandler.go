package view

import (
	"golang_beginning/hw14/models"
	"html/template"
	"net/http"
)

var prodListTemplate = template.Must(template.ParseFiles("hw14/view/items.html"))

func WithTemplate(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		prodListTemplate.Execute(w, models.Items)
	}
}
