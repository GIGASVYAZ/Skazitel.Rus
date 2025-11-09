package messagecommand

import "context"

type SendMessageCommand struct {
	UserID  int64
	Content string
}

type SendMessageHandler struct {
	messageRepo MessageRepository
}

func NewSendMessageHandler(messageRepo MessageRepository) *SendMessageHandler {
	return &SendMessageHandler{
		messageRepo: messageRepo,
	}
}

func (h *SendMessageHandler) Handle(ctx context.Context, cmd SendMessageCommand) error {
	return h.messageRepo.Create(cmd.UserID, cmd.Content)
}
