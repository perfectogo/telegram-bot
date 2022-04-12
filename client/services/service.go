package services

import (
	"fmt"

	"github.com/perfectogo/telegram-bot/config"
	pb "github.com/perfectogo/telegram-bot/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type InterfaceServiceManager interface {
	TgService() pb.MessageSenderClient
}
type serviceManager struct {
	tgService pb.MessageSenderClient
}

func (s *serviceManager) TgService() pb.MessageSenderClient {
	return s.tgService
}

func NewServiceManager(cfg *config.Config) (InterfaceServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	conntg, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ServiceHost, cfg.ServicePORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		tgService: pb.NewMessageSenderClient(conntg),
	}

	return serviceManager, nil
}
