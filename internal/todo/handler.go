package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/utils"
)

var todos = []Todo{}
var nextID = 1

// GetTodos godoc
// @Summary      Get all todos
// @Tags         todos
// @Produce      json
// @Success      200  {array}  Todo
// @Router       /todos [get]
func getTodos(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo godoc
// @Summary      Create a new todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo body Todo true "Todo to create"
// @Success      200  {object}  Todo
// @Router       /todos [post]
func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	if err := utils.DecodeStrictJson(r, &t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if t.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	t.ID = nextID
	nextID++
	todos = append(todos, t)
	json.NewEncoder(w).Encode(t)
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// UpdateTodo godoc
// @Summary     Update a todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id   path     int   true  "Todo ID"
// @Param       todo body     Todo  true  "Updated todo"
// @Success     200  {object} Todo
// @Failure     400  {string} string "Bad Request"
// @Failure     404  {string} string "Not Found"
// @Router      /todos/{id} [put]
func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromPath(r, "/todos/")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updated.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
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

	http.NotFound(w, r)
}

// DeleteTodo godoc
// @Summary     Delete a todo
// @Tags        todos
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success     204  {string} string "No Content"
// @Failure     404  {string} string "Not Found"
// @Router      /todos/{id} [delete]
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromPath(r, "/todos/")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func TodoItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
