package unittest

import (
	"context"
	"fmt"
	"os"
	"skazitel-rus/internal/domain/message"
	"skazitel-rus/internal/domain/user"
	"skazitel-rus/pkg/config"
	"skazitel-rus/pkg/database"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

var pool *pgxpool.Pool
var testUserId *int64

func TestMain(m *testing.M) {
	fmt.Println("Тесты запущены")

	cfg := config.New()
	ctx := context.Background()

	database.InitPoolWithConfig(
		ctx,
		cfg.Database.URL,
		cfg.Database.MaxConns,
		cfg.Database.MinConns,
	)

	pool = database.GetPool()
	if pool == nil {
		fmt.Println("Ошибка: пул подключений не инициализирован")
		os.Exit(1)
	}

	exitVal := m.Run()

	fmt.Println("Очистка ресурсов")
	database.ClosePool()

	os.Exit(exitVal)
}

func setupDB(t *testing.T) {
	ctx := context.Background()

	require.NotNil(t, pool, "пул подключений не инициализирован")

	_, err := pool.Exec(ctx, "drop schema if exists skazitel cascade; create schema skazitel;")
	require.NoError(t, err, "ошибка при инициализации БД")

	_, err = pool.Exec(ctx, user.UserTableSQL)
	require.NoError(t, err, "ошибка при инициализации БД")

	_, err = pool.Exec(ctx, message.MessageTableSQL)
	require.NoError(t, err, "ошибка при инициализации БД")

	_, err = pool.Exec(ctx, `
    insert into skazitel.users (username, password)
    values ('test-user', 'test-pass');
  `)
	require.NoError(t, err, "ошибка при инициализации БД")

	err = pool.QueryRow(ctx, "select id from skazitel.users where username = 'test-user';").Scan(&testUserId)
	require.NoError(t, err, "ошибка при инициализации БД")
}
