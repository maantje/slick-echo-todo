package main

import (
	"log"
	"todoist/handler"
	"todoist/model"

	"github.com/anthdm/slick"
)

func main() {
	app := slick.New()

	todoHandler := handler.TodoHandler{Todos: make([]model.Todo, 0)}

	app.Get("/", todoHandler.Index)
	app.Post("/todo", todoHandler.Store)
	app.Post("/todo/:id/delete", todoHandler.Destroy)
	app.Post("/todo/:id/toggle", todoHandler.Toggle)

	log.Fatal(app.Start())
}
