package chat

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

const DEF_MESSAGES = `CREATE TABLE skazitel.messages (
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES skazitel.users(id),
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);`

type Message struct {
	Id        int64
	UsersId   int64
	Content   string
	CreatedAt time.Time
}

func CreateMessage(UserId int64, Content string) error {
	var err error
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(),
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, $2)", UserId, Content)
	if err != nil {
		return err
	}

	return nil
}

func GetMessages(limit int) ([]Message, error) {
	var err error
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(
		context.Background(),
		`SELECT id, user_id, content, created_at 
         FROM (
             SELECT id, user_id, content, created_at 
             FROM skazitel.messages 
             ORDER BY created_at DESC 
             LIMIT $1
         ) AS last_messages
         ORDER BY created_at ASC`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.Id, &msg.UsersId, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
