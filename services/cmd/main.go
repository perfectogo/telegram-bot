package main

import (
	"fmt"
	"net"

	"github.com/perfectogo/telegram-bot/config"
	pb "github.com/perfectogo/telegram-bot/genproto"
	"github.com/perfectogo/telegram-bot/pkg/logger"
	"github.com/perfectogo/telegram-bot/services/service"
	"google.golang.org/grpc"
)

func main() {
	cfg, _ := config.Load()

	fmt.Println(cfg)
	// Loading logger
	log := logger.New(cfg.LogLevel, "catalog-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	lis, err := net.Listen("tcp", cfg.ServicePORT)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	// goroutine for ...
	go service.Sender()

	tgService := service.NewTgService(log)
	newService := grpc.NewServer()
	pb.RegisterMessageSenderServer(newService, tgService)
	log.Info("main: server running",
		logger.String("port", cfg.ServicePORT))

	if err := newService.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
