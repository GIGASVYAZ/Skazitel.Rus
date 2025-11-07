package handler

import (
	"net/http"
	"skazitel-rus/internal/usecase"
	"skazitel-rus/pkg/response"
	"strconv"
)

type SendMessageRequest struct {
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
}

type Handler struct {
	messageUseCase *usecase.MessageUseCase
	userUseCase    *usecase.UserUseCase
}

func NewHandler(messageUC *usecase.MessageUseCase, userUC *usecase.UserUseCase) *Handler {
	return &Handler{
		messageUseCase: messageUC,
		userUseCase:    userUC,
	}
}

func (h *Handler) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.SendMessageHandler(w, r)
	} else if r.Method == http.MethodGet {
		h.GetMessagesHandler(w, r)
	} else {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
	}
}

func (h *Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.ErrorResponse(w, http.StatusNotFound, "Маршрут не найден")
}

func (h *Handler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	var req SendMessageRequest
	err := response.DecodeJSON(r, &req)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	err = h.messageUseCase.SendMessage(req.UserId, req.Content)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при создании сообщения")
		return
	}

	response.SuccessResponse(w, http.StatusCreated, nil, "Сообщение отправлено")
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
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

	messages, err := h.messageUseCase.GetLastMessages(limit)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при получении сообщений")
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages, "Сообщения получены")
}
