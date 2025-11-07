package usecase

import (
	"skazitel-rus/internal/domain"
	"skazitel-rus/internal/repository"
)

type MessageUseCase struct {
	messageRepo repository.MessageRepository
}

func NewMessageUseCase(messageRepo repository.MessageRepository) *MessageUseCase {
	return &MessageUseCase{
		messageRepo: messageRepo,
	}
}

func (uc *MessageUseCase) SendMessage(userId int64, content string) error {
	return uc.messageRepo.Create(userId, content)
}

func (uc *MessageUseCase) GetLastMessages(limit int) ([]domain.Message, error) {
	return uc.messageRepo.GetNLast(limit)
}
