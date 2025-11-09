package message

import "time"

type Message struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

type MessageRepository interface {
	Create(userID int64, content string) error
	GetLastN(limit int) ([]Message, error)
}

const MessageTableSQL = `CREATE TABLE IF NOT EXISTS skazitel.messages (
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES skazitel.users(id),
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);`
