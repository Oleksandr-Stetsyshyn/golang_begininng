package translator

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/translate", translationHandler)
	http.ListenAndServe(":8080", r)

}
