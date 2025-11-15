package getmessageusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
	"strconv"
)

func (h *GetMessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
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

	q := GetMessagesQuery{
		Limit: limit,
	}

	messages, err := h.Handle(r.Context(), q)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка при получении сообщений")
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages, "Сообщения получены")
}
