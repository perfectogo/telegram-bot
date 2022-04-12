package handlers

import (
	"context"

	pb "github.com/perfectogo/telegram-bot/genproto"
	"github.com/perfectogo/telegram-bot/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (h *handler) SendMessageHandler(ctx *gin.Context) {

	var newMessage pb.MessageRequest

	if err := ctx.BindJSON(&newMessage); err != nil {
		ctx.AbortWithStatusJSON(500, "failed to bind json")
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}
	res, err := h.serviceManager.TgService().Sender(context.Background(), &pb.MessageRequest{
		Text:     newMessage.Text,
		Priority: newMessage.Priority,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(500, "fieald sending message")
		h.log.Error("failed sending message", logger.Error(err))
		return
	}
	ctx.JSON(201, res)
}
