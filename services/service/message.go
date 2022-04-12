package service

import (
	"context"
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
		s.logger.Info("Undefined Priority")
	}

	res := &pb.MessageResponse{
		Message: req.Text,
	}

	return res, nil
}

func Sender() {

	for {
		if len(HighPriorityMessages) > 0 {
			err := bot.MessageSenderBot(HighPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("couldn'n send message to tg_bot: %v", err)
			}
			HighPriorityMessages = Remove(HighPriorityMessages, 0)

		} else if len(MediumPriorityMessages) > 0 {
			err := bot.MessageSenderBot(MediumPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("couldn'n send message to tg_bot: %v", err)
			}
			MediumPriorityMessages = Remove(MediumPriorityMessages, 0)

		} else if len(LowPriorityMessages) > 0 {
			err := bot.MessageSenderBot(LowPriorityMessages[0].Text)
			if err != nil {
				log.Fatalf("couldn'n send message to tg_bot: %v", err)
			}
			LowPriorityMessages = Remove(LowPriorityMessages, 0)
		}
		time.Sleep(time.Second * 5)
	}
}

func Remove(mReq []pb.MessageRequest, i int) []pb.MessageRequest {
	mReq[i] = mReq[len(mReq)-1]
	return mReq[:len(mReq)-1]
}
