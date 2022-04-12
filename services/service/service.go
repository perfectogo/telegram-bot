package service

import "github.com/perfectogo/telegram-bot/pkg/logger"

type tgService struct {
	logger logger.Logger
}

func NewTgService(log logger.Logger) *tgService {
	return &tgService{
		logger: log,
	}
}
