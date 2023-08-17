package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{1, "Task 1"},
	{2, "Task 2"},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func main() {
	http.HandleFunc("/tasks", getTasks)

	fmt.Println("server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
