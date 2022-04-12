package service

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/perfectogo/telegram-bot/bot"
	pb "github.com/perfectogo/telegram-bot/genproto"
)

var HighPriorityMessages []pb.MessageRequest
var MediumPriorityMessages []pb.MessageRequest
var LowPriorityMessages []pb.MessageRequest

func (s *tgService) Sender(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	fmt.Println("Sender function was invoked")

	var message pb.MessageRequest

	message.Text = req.Text
	message.Priority = req.Priority

	if message.Priority == "high" {
		HighPriorityMessages = append(HighPriorityMessages, message)
	} else if message.Priority == "medium" {
		MediumPriorityMessages = append(MediumPriorityMessages, message)
	} else if message.Priority == "low" {
		LowPriorityMessages = append(LowPriorityMessages, message)
	} else {
		fmt.Println("Undefined Priority")
	}

	res := &pb.MessageResponse{
		Message: req.Text,
	}
	fmt.Println(message.Text)

	return res, nil
}

func Sender() {

	for {
		if len(HighPriorityMessages) > 0 {
			err := bot.MessageSenderBot(HighPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			}
			HighPriorityMessages = Remove(HighPriorityMessages, 0)
		} else if len(MediumPriorityMessages) > 0 {
			err := bot.MessageSenderBot(MediumPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			}
			MediumPriorityMessages = Remove(MediumPriorityMessages, 0)
		} else if len(LowPriorityMessages) > 0 {
			err := bot.MessageSenderBot(LowPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("Problem with sending message to bot: %v", err)
			}
			LowPriorityMessages = Remove(LowPriorityMessages, 0)
		}
		time.Sleep(time.Second * 10)
	}
}

func Remove(s []pb.MessageRequest, i int) []pb.MessageRequest {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
