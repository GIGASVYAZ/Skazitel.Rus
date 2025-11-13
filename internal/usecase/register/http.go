package registerusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
)

func (h *RegisterUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	err := response.DecodeJSON(r, &req)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	cmd := RegisterUserCommand{
		Username: req.Username,
		Password: req.Password,
	}

	err = h.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка регистрации")
		return
	}

	response.SuccessResponse(w, http.StatusCreated, nil, "Пользователь зарегистрирован")
}
