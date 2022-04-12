package handlers

import (
	"github.com/perfectogo/telegram-bot/client/services"
	"github.com/perfectogo/telegram-bot/config"
	"github.com/perfectogo/telegram-bot/package/logger"
)

type handler struct {
	log            logger.Logger
	serviceManager services.InterfaceServiceManager
	cfg            config.Config
}

// HandlerConfig ...
type HandlerConfig struct {
	Logger         logger.Logger
	ServiceManager services.InterfaceServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerConfig) *handler {
	return &handler{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
