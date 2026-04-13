package main

import (
	"log"
	"rest-api/http"
	"rest-api/todo"
)

func main() {
	todolist := todo.NewList()

	httphandlers := http.NewHTTPHandlers(todolist)
	httpserver := http.NewHTTPServer(httphandlers)

	if err := httpserver.StartServer(); err != nil {
		log.Println(err.Error())
	}
}