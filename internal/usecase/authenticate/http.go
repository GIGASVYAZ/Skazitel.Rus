package authenticateusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
)

func (h *AuthenticateUserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
		return
	}

	type AuthenticateRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req AuthenticateRequest
	err := response.DecodeJSON(r, &req)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Неверный формат данных")
		return
	}

	q := AuthenticateUserQuery(req)

	result, err := h.Handle(r.Context(), q)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, result, "Аутентификация успешна")
}
