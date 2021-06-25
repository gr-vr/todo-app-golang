package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/satanCoffee/todo-app-golang"
	"github.com/satanCoffee/todo-app-golang/pkg/handler"
	"github.com/satanCoffee/todo-app-golang/pkg/repository"
	"github.com/satanCoffee/todo-app-golang/pkg/service"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
