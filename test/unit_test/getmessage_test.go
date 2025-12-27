package unittest

import (
	"context"
	"fmt"
	"os"
	"skazitel-rus/internal/domain/message"
	"skazitel-rus/internal/domain/user"
	messagerepository "skazitel-rus/internal/infrastructure/repository/message"
	getmessageusecase "skazitel-rus/internal/usecase/getmessage"
	"skazitel-rus/pkg/config"
	"skazitel-rus/pkg/database"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
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

	_, err := pool.Exec(ctx, "drop schema if exists skazitel cascade;")
	pool.Exec(ctx, user.UserTableSQL)
	pool.Exec(ctx, message.MessageTableSQL)
	pool.Exec(ctx, `
		insert into skazitel.users (username, password)
		values ('test-user', 'test-pass');
	`)
	pool.QueryRow(ctx, "select id from skazitel.users where username = 'test-user';").Scan(testUserId)
	require.NoError(t, err, "ошибка при инициализации БД")
}

func setup(t *testing.T) *getmessageusecase.GetMessagesHandler {
	setupDB(t)
	repo := messagerepository.New(pool)
	handler := getmessageusecase.NewGetMessagesHandler(repo)
	return handler
}

func TestNoneMessagesReturnsEmptyList(t *testing.T) {
	ctx := context.Background()
	handler := setup(t)

	result, err := handler.Handle(ctx, getmessageusecase.GetMessagesQuery{
		Limit: 2,
	})

	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, len(result))
}

func TestReturnsMax2Messages(t *testing.T) {
	ctx := context.Background()
	handler := setup(t)

	_, err := pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES (23, 'first')")

	_, err = pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES (23, 'second')")

	_, err = pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES (23, 'third')")
	require.NoError(t, err, "ошибка при вставке тестовых данных")

	result, err := handler.Handle(ctx, getmessageusecase.GetMessagesQuery{
		Limit: 2,
	})

	assert.NoError(t, err, "ошибка при вызове обработчика")
	assert.NotEmpty(t, result, "результат не должен быть пустым")
	assert.Equal(t, 2, len(result), "должно быть ровно 2 сообщения")
	assert.Equal(t, "second", result[0].Content)
	assert.Equal(t, "third", result[1].Content)
}
