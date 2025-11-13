package message

import "time"

type Message struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageRepository interface {
	SendMessage(userID int64, content string) error
	GetNLast(limit int) ([]Message, error)
}

const MessageTableSQL = `CREATE TABLE IF NOT EXISTS skazitel.messages (
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES skazitel.users(id),
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);`
