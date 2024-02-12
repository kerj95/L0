package main

import (
	"L0"
	"L0/pkg/handler"
	"L0/pkg/repository"
	"L0/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//go natsSubscriber.NatsSubscriber()

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(L0.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	//log.Print("Server Started")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
