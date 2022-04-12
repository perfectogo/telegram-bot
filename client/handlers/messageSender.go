package handlers

import (
	"context"
	"fmt"
	"log"

	pb "github.com/perfectogo/telegram-bot/genproto"

	"github.com/gin-gonic/gin"
)

var c pb.MessageSenderClient

func (h *handler) SendMessageHandler(ctx *gin.Context) {

	var newMessage pb.MessageRequest

	if err := ctx.BindJSON(&newMessage); err != nil {

		log.Printf("Failed to accept message from cleint")
	}
	fmt.Println(newMessage)
	res, err := c.Sender(context.Background(), &pb.MessageRequest{
		Text:     newMessage.Text,
		Priority: newMessage.Priority,
	})

	if err != nil {

		log.Fatalf("Failed to call Sender RPC: %v", err)
	}
	log.Println(res)
}
