package handler

import (
	"net/http"
	"strconv"

	messagecommand "skazitel-rus/internal/usecase/command/message"
	messagequery "skazitel-rus/internal/usecase/query/message"
	"skazitel-rus/pkg/response"
)

type MessageHandler struct {
	sendMessageHandler *messagecommand.SendMessageHandler
	getMessagesHandler *messagequery.GetMessagesHandler
}

func NewMessageHandler(
	sendMessageHandler *messagecommand.SendMessageHandler,
	getMessagesHandler *messagequery.GetMessagesHandler,
) *MessageHandler {
	return &MessageHandler{
		sendMessageHandler: sendMessageHandler,
		getMessagesHandler: getMessagesHandler,
	}
}

func (h *MessageHandler) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.SendMessage(w, r)
	case http.MethodGet:
		h.GetMessages(w, r)
	default:
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
	}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	type SendMessageRequest struct {
		UserID  int64  `json:"user_id"`
		Content string `json:"content"`
	}

	var req SendMessageRequest
	err := response.DecodeJSON(r, &req)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	cmd := messagecommand.SendMessageCommand{
		UserID:  req.UserID,
		Content: req.Content,
	}

	err = h.sendMessageHandler.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при создании сообщения")
		return
	}

	response.SuccessResponse(w, http.StatusCreated, nil, "Сообщение отправлено")
}

func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, "Неверный параметр limit")
			return
		}
	}

	q := messagequery.GetMessagesQuery{
		Limit: limit,
	}

	messages, err := h.getMessagesHandler.Handle(r.Context(), q)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при получении сообщений")
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages, "Сообщения получены")
}
