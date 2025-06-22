package main

import (
	"log"
	"net/http"

	"todo-api/todo"
)

func main() {
	http.HandleFunc("/todos", todo.TodoHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
