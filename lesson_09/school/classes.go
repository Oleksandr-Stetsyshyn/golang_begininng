package school

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var A = Class{
	Name: "4A",
	Students: []Student{
		{Name: "John Doe", averageGrade: 85.5},
		{Name: "Jane Smith", averageGrade: 92.3},
		{Name: "Bob Johnson", averageGrade: 78.9},
		{Name: "Alice Lee", averageGrade: 91.2},
		{Name: "Mike Brown", averageGrade: 87.4},
		{Name: "Emily Davis", averageGrade: 89.1},
		{Name: "David Wilson", averageGrade: 83.7},
		{Name: "Sarah Taylor", averageGrade: 94.5},
		{Name: "Kevin Chen", averageGrade: 88.2},
	},
}

var B = Class{
	Name: "2B",
	Students: []Student{
		{Name: "John Doe", averageGrade: 85.5},
		{Name: "Jane Smith", averageGrade: 92.3},
		{Name: "Bob Johnson", averageGrade: 78.9},
		{Name: "Alice Lee", averageGrade: 91.2},
		{Name: "Mike Brown", averageGrade: 87.4},
		{Name: "Emily Davis", averageGrade: 89.1},
		{Name: "David Wilson", averageGrade: 83.7},
		{Name: "Sarah Taylor", averageGrade: 94.5},
		{Name: "Kevin Chen", averageGrade: 88.2},
	},
}

var School = SchoolInfo{
	Name:    "School #1",
	Classes: []Class{A, B},
}

var Classes http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	schoolInfo := map[string]interface{}{
		"name":       School.Name,
		"numClasses": len(School.Classes),
	}
	response, err := json.MarshalIndent(schoolInfo, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

var ClassInfo http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	if vars["name"] == A.Name {
		response, err := json.MarshalIndent(A, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}
	if vars["name"] == B.Name {
		response, err := json.MarshalIndent(B, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}
	w.Write([]byte("[]"))
}
