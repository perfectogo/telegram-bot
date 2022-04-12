package main

import (
	"github.com/perfectogo/telegram-bot/client/routes"
	"github.com/perfectogo/telegram-bot/client/services"
	"github.com/perfectogo/telegram-bot/config"
	"github.com/perfectogo/telegram-bot/package/logger"
)

func main() {
	cfg, _ := config.Load()

	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := routes.New(routes.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.ClientPORT); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
