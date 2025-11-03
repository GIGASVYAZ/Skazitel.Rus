package repository

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain"
	"skazitel-rus/pkg/database"
	"time"
)

func CreateMessage(UserId int64, Content string) error {
	pool := database.GetPool()
	if pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, $2)",
		UserId, Content)

	return err
}

func GetMessages(limit int) ([]domain.Message, error) {
	pool := database.GetPool()
	if pool == nil {
		return nil, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := pool.Query(ctx,
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

	var messages []domain.Message
	for rows.Next() {
		var msg domain.Message
		err := rows.Scan(&msg.Id, &msg.UsersId, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
