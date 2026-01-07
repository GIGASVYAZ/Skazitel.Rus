package unittest

import (
	"context"
	messagerepository "skazitel-rus/internal/infrastructure/repository/message"
	sendmessageusecase "skazitel-rus/internal/usecase/sendmessage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupSendMessage(t *testing.T) *sendmessageusecase.SendMessageHandler {
	setupDB(t)
	repo := messagerepository.New(pool)
	handler := sendmessageusecase.NewSendMessageHandler(repo)
	return handler
}
func TestSendMessageSuccess(t *testing.T) {
	ctx := context.Background()
	handler := setupSendMessage(t)

	messageContent := "Hello, World!"
	cmd := sendmessageusecase.SendMessageCommand{
		UserID:  *testUserId,
		Content: messageContent,
	}

	err := handler.Handle(ctx, cmd)

	assert.NoError(t, err, "Handle не должен возвращать ошибку")

	var count int
	err = pool.QueryRow(ctx,
		"SELECT count(*) FROM skazitel.messages WHERE user_id = $1 AND content = $2",
		*testUserId, messageContent).Scan(&count)

	require.NoError(t, err, "Ошибка при проверке данных в БД")
	assert.Equal(t, 1, count, "Сообщение должно быть сохранено в БД")
}

func TestSendMessageToNonExistentUser(t *testing.T) {
	ctx := context.Background()
	handler := setupSendMessage(t)

	nonExistentUserID := int64(999999)
	cmd := sendmessageusecase.SendMessageCommand{
		UserID:  nonExistentUserID,
		Content: "Ghost message",
	}

	err := handler.Handle(ctx, cmd)

	assert.Error(t, err, "Должна быть ошибка из-за несуществующего пользователя")
}
