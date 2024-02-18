package main

import (
	"log"

	"github.com/RapchynskyiOleksandr/go-todo-list"
	"github.com/RapchynskyiOleksandr/go-todo-list/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Errorwith running http server: %s", err.Error())
	}
}
