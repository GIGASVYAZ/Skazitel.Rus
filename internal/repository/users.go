package repository

import (
	"context"
	"errors"
	"skazitel-rus/pkg/database"
	"time"
)

func CreateUser(Username string, Password string) error {
	pool := database.GetPool()
	if pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pool.Exec(ctx,
		"INSERT INTO skazitel.users (username, password) VALUES ($1, $2)",
		Username, Password)

	return err
}

func UserAuthenticate(Username string, Password string) (bool, error) {
	pool := database.GetPool()
	if pool == nil {
		return false, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var passwordUser string
	err := pool.QueryRow(ctx,
		"SELECT password FROM skazitel.users WHERE username = $1",
		Username).Scan(&passwordUser)

	if err != nil {
		return false, err
	}

	if Password == passwordUser {
		return true, nil
	}

	return false, errors.New("пароль неверный")
}

func UpdateUserStatus(Username string, IsOnline bool) error {
	pool := database.GetPool()
	if pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tag, err := pool.Exec(ctx,
		"UPDATE skazitel.users SET is_online = $2 WHERE username = $1",
		Username, IsOnline)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("пользователь не найден")
	}

	return nil
}
