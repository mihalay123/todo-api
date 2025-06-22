package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/internal/utils"
)

var todos = []Todo{}
var nextID = 1

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var t Todo
		if err := utils.DecodeStrictJson(r, &t); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		t.ID = nextID
		nextID++
		todos = append(todos, t)
		json.NewEncoder(w).Encode(t)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TodoItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var updated Todo
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Title = updated.Title
				todos[i].Done = updated.Done
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		}

	case http.MethodDelete:
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.NotFound(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
