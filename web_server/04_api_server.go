// * Web Service API Server
// studi kasus task manager (seperti yang kemarin CLI)
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID          string
	Name        string
	IsCompleted bool
}

type Tasks []Task

var tasksData = Tasks{
	Task{"UU001", "mencuci piring", true},
	Task{"UU002", "belajar go", false},
	Task{"UU003", "masak nasgor", false},
	Task{"UU004", "mandi malem", false},
}

// * GET /tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	// set response header
	w.Header().Set("Content-Type", "application/json")

	// guard for method is not GET
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// convert object to json
	result, err := json.Marshal(tasksData)
	if err != nil {
		// handle if any error on converting
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response
	w.Write(result)
	return
}

// * GET /task by id
func getTask(w http.ResponseWriter, r *http.Request) {
	// set response header
	w.Header().Set("Content-Type", "application/json")

	// guard for method is not GET
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// get form value of id
	taskId := r.FormValue("id")

	// find data on object that equal input id
	for _, task := range tasksData {
		if task.ID == taskId {

			taskJson, err := json.Marshal(task)
			if err != nil {
				// handle if any error on converting
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// write response
			w.Write(taskJson)
			return
		}
	}

	// error when data not found
	http.Error(w, "Task not found", http.StatusNotFound)
}

// * Main Func
func main() {
	http.HandleFunc("/tasks", getAllTasks)
	http.HandleFunc("/task", getTask)

	fmt.Println("Start server at :8080...")
	http.ListenAndServe(":8080", nil)
}
