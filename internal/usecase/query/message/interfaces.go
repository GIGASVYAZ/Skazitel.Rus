package messagequery

import "skazitel-rus/internal/domain/message"

type MessageRepository interface {
	GetLastN(limit int) ([]message.Message, error)
}
