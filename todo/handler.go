package todo

import (
	"encoding/json"
	"net/http"
)

var todos = []Todo{}
var nextID = 1

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var t Todo
		json.NewDecoder(r.Body).Decode(&t)
		t.ID = nextID
		nextID++
		todos = append(todos, t)
		json.NewEncoder(w).Encode(t)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
