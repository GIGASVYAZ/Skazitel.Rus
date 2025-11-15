package sendmessageusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
)

func (h *SendMessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
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

	cmd := SendMessageCommand{
		UserID:  req.UserID,
		Content: req.Content,
	}

	err = h.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при создании сообщения")
		return
	}

	response.SuccessResponse(w, http.StatusCreated, nil, "Сообщение отправлено")
}
