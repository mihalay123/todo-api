package main

import (
	"log"
	"net/http"

	_ "todo-api/docs"
	"todo-api/internal/todo"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           ToDo API
// @version         1.0
// @description     Простое API для управления задачами
// @host            localhost:8080
// @BasePath        /

func main() {
	http.Handle("/docs/", httpSwagger.WrapHandler)
	http.HandleFunc("/todos", todo.TodoHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
