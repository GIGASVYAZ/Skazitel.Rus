package setonlineusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
)

func (h *SetUserOnlineHandler) SetOnline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	type SetOnlineRequest struct {
		Username string `json:"username"`
		IsOnline bool   `json:"is_online"`
	}

	var req SetOnlineRequest
	err := response.DecodeJSON(r, &req)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	cmd := SetUserOnlineCommand(req)

	err = h.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, nil, "Статус обновлен")
}
