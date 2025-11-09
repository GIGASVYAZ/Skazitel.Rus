package userrepository

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain/user"
	"skazitel-rus/pkg/database"
	"time"
)

type PostgresUserRepository struct{}

func New() *PostgresUserRepository {
	return &PostgresUserRepository{}
}

func (r *PostgresUserRepository) Create(username string, password string) error {
	pool := database.GetPool()
	if pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx,
		"INSERT INTO skazitel.users (username, password) VALUES ($1, $2)",
		username, password)
	return err
}

func (r *PostgresUserRepository) GetByUsername(username string) (*user.User, error) {
	pool := database.GetPool()
	if pool == nil {
		return nil, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &user.User{}
	err := pool.QueryRow(ctx,
		"SELECT id, username, password, created_at, is_online FROM skazitel.users WHERE username = $1",
		username).Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.IsOnline)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresUserRepository) UpdateIsOnline(username string, isOnline bool) error {
	pool := database.GetPool()
	if pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tag, err := pool.Exec(ctx,
		"UPDATE skazitel.users SET is_online = $2 WHERE username = $1",
		username, isOnline)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("пользователь не найден")
	}

	return nil
}
