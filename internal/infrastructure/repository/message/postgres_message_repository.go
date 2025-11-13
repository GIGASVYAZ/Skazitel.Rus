package messagerepository

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain/message"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresMessageRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PostgresMessageRepository {
	return &PostgresMessageRepository{pool: pool}
}

func (r *PostgresMessageRepository) SendMessage(userID int64, content string) error {
	if r.pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, $2)",
		userID, content)
	return err
}

func (r *PostgresMessageRepository) GetNLast(limit int) ([]message.Message, error) {
	if r.pool == nil {
		return nil, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, content, created_at
		FROM (
			SELECT id, user_id, content, created_at
			FROM skazitel.messages
			ORDER BY created_at DESC
			LIMIT $1
		) AS last_messages
		ORDER BY created_at ASC
	`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []message.Message
	for rows.Next() {
		var msg message.Message
		err := rows.Scan(&msg.ID, &msg.UserID, &msg.Content, &msg.CreatedAt)
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
