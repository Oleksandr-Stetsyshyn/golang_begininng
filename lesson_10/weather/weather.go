package weather

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", cityWeather)
	http.ListenAndServe(":8080", r)

}
