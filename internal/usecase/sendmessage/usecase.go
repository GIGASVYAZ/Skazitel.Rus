package sendmessageusecase

import (
	"context"
	"skazitel-rus/internal/domain/message"
)

type SendMessageCommand struct {
	UserID  int64
	Content string
}

type SendMessageHandler struct {
	messageRepo message.MessageRepository
}

func NewSendMessageHandler(messageRepo message.MessageRepository) *SendMessageHandler {
	return &SendMessageHandler{
		messageRepo: messageRepo,
	}
}

func (h *SendMessageHandler) Handle(ctx context.Context, cmd SendMessageCommand) error {
	return h.messageRepo.SendMessage(cmd.UserID, cmd.Content)
}
