package todo

import (
	"encoding/json"
	"net/http"
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
