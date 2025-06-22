package main

import (
	"log"
	"net/http"

	_ "todo-api/docs"
	"todo-api/internal/todo"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Todo API
// @version         1.0
// @description     A simple API to manage a todo list in memory

// @contact.name   Mike
// @contact.email  mihalay26@gmail.com

// @host      localhost:8080
// @BasePath  /

func main() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/todos", todo.TodoHandler)
	http.HandleFunc("/todos/", todo.TodoItemHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
