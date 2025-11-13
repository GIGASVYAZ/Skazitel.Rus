package userrepository

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain/user"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) RegisterUser(username string, password string) error {
	if r.pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.pool.Exec(ctx,
		"INSERT INTO skazitel.users (username, password) VALUES ($1, $2)",
		username, password)
	return err
}

func (r *PostgresUserRepository) GetByUsername(username string) (*user.User, error) {
	if r.pool == nil {
		return nil, errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &user.User{}
	err := r.pool.QueryRow(ctx,
		"SELECT id, username, password, created_at, is_online FROM skazitel.users WHERE username = $1",
		username).Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.IsOnline)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}

func (r *PostgresUserRepository) UpdateIsOnline(username string, isOnline bool) error {
	if r.pool == nil {
		return errors.New("пул подключений не инициализирован")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tag, err := r.pool.Exec(ctx,
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
