package main

import (
	"exam3/api"
	"exam3/config"
	"exam3/pkg/logger"
	"exam3/service"
	"exam3/storage/postgres"
	"fmt"
	"golang.org/x/net/context"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	pgStore, err := postgres.New(context.Background(), cfg, log)
	if err != nil {
		log.Error("error is while connecting to db", logger.Error(err))
		return
	}
	defer pgStore.Close()

	services := service.New(pgStore, log)

	server := api.New(services, log)

	fmt.Println(server)

	log.Info("Service is running on", logger.Int("port", 8090))
	if err := server.Run("localhost:8090"); err != nil {
		panic(err)
	}
}
