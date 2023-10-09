package todoList

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type daylyTasks struct {
	Date  string   `json:"date"`
	Tasks []string `json:"tasks"`
}

func Main() {
	tasks := []daylyTasks{
		{
			Date:  "2022-01-01",
			Tasks: []string{"Task 1", "Task 2", "Task 3"},
		},
		{
			Date:  "2022-01-02",
			Tasks: []string{"Task 4", "Task 5"},
		},
		{
			Date:  "2022-01-03",
			Tasks: []string{"Task 6", "Task 7", "Task 8", "Task 9"},
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	r.HandleFunc("/tasks/{date}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		vars := mux.Vars(r)
		for _, task := range tasks {
			if task.Date == vars["date"] {
				response, err := json.MarshalIndent(task, "", "  ")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(response)
				return
			}
		}
		w.Write([]byte("[]"))
	})

	http.ListenAndServe(":8080", r)

}
