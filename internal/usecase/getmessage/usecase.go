package getmessageusecase

import (
	"context"
	"skazitel-rus/internal/domain/message"
)

type GetMessagesQuery struct {
	Limit int
}

type GetMessagesHandler struct {
	messageRepo message.MessageRepository
}

func NewGetMessagesHandler(messageRepo message.MessageRepository) *GetMessagesHandler {
	return &GetMessagesHandler{
		messageRepo: messageRepo,
	}
}

func (h *GetMessagesHandler) Handle(ctx context.Context, q GetMessagesQuery) ([]MessageDTO, error) {
	messages, err := h.messageRepo.GetNLast(q.Limit)
	if err != nil {
		return nil, err
	}

	dtos := make([]MessageDTO, len(messages))
	for i, msg := range messages {
		dtos[i] = MessageDTO{
			ID:        msg.ID,
			UserID:    msg.UserID,
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt,
		}
	}

	return dtos, nil
}
