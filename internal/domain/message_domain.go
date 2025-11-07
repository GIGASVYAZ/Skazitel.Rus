package domain

import "time"

type Message struct {
	Id        int64
	UsersId   int64
	Content   string
	CreatedAt time.Time
}

const MessageTableSQL = `CREATE TABLE IF NOT EXISTS skazitel.messages (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES skazitel.users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);`
