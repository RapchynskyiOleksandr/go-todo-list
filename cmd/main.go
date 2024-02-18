package main

import (
	"log"

	"github.com/RapchynskyiOleksandr/go-todo-list"
	"github.com/RapchynskyiOleksandr/go-todo-list/pkg/handler"
	"github.com/RapchynskyiOleksandr/go-todo-list/pkg/repository"
	"github.com/RapchynskyiOleksandr/go-todo-list/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Errorwith running http server: %s", err.Error())
	}
}
