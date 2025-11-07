package repository

import (
	"skazitel-rus/internal/domain"
)

type UserRepository interface {
	Create(username string, password string) error
	IsPasswordByUsernameEqualTo(username string, password string) (bool, error)
	UpdateIsOnline(username string, isOnline bool) error
}

type MessageRepository interface {
	Create(userId int64, content string) error
	GetNLast(limit int) ([]domain.Message, error)
}
