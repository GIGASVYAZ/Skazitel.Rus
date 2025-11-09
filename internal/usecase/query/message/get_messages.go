package messagequery

import (
	"context"
)

type GetMessagesQuery struct {
	Limit int
}

type GetMessagesHandler struct {
	messageRepo MessageRepository
}

func NewGetMessagesHandler(messageRepo MessageRepository) *GetMessagesHandler {
	return &GetMessagesHandler{
		messageRepo: messageRepo,
	}
}

func (h *GetMessagesHandler) Handle(ctx context.Context, q GetMessagesQuery) ([]MessageDTO, error) {
	messages, err := h.messageRepo.GetLastN(q.Limit)
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
