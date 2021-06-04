package main

import (
	"github.com/satanCoffee/todo-app-golang"
	"github.com/satanCoffee/todo-app-golang/pkg/handler"
	"github.com/satanCoffee/todo-app-golang/pkg/repository"
	"github.com/satanCoffee/todo-app-golang/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main()  {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}