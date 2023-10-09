package school

import (
	"github.com/gorilla/mux"
	"net/http"
)

type SchoolInfo struct {
	Name    string  `json:"name"`
	Classes []Class `json:"classes"`
}

type Student struct {
	Name         string  `json:"name"`
	averageGrade float64 `json:"averageGrade"`
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

func (c Class) getAverageGrade() float64 {
	var sum float64
	for _, student := range c.Students {
		sum += student.averageGrade
	}
	return sum / float64(len(c.Students))
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/login", ApplyMiddleware(signIn, Authenticate)).Methods("POST")

	r.HandleFunc("/class", Classes)
	r.HandleFunc("/class/{name}", Authorize(ClassInfo))

	http.ListenAndServe("localhost:8080", r)
}
