package unittest

import (
	"context"
	messagerepository "skazitel-rus/internal/infrastructure/repository/message"
	getmessageusecase "skazitel-rus/internal/usecase/getmessage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, 'first')", testUserId)

	_, err = pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, 'second')", testUserId)

	_, err = pool.Exec(ctx,
		"INSERT INTO skazitel.messages (user_id, content) VALUES ($1, 'third')", testUserId)
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
