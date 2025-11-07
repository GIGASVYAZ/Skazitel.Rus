package userRepository

import (
	"context"
	"errors"
	"skazitel-rus/pkg/database"
	"time"
)

type PostgresUserRepository struct{}

func NewPostgresUserRepository() *PostgresUserRepository {
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

func (r *PostgresUserRepository) IsPasswordByUsernameEqualTo(username string, password string) (bool, error) {
	pool := database.GetPool()
	if pool == nil {
		return false, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var passwordUser string
	err := pool.QueryRow(ctx,
		"SELECT password FROM skazitel.users WHERE username = $1",
		username).Scan(&passwordUser)

	if err != nil {
		return false, err
	}

	if password == passwordUser {
		return true, nil
	}

	return false, errors.New("пароль неверный")
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
