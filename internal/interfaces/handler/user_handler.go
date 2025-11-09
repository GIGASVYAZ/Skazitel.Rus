package handler

import (
	"net/http"

	usercommand "skazitel-rus/internal/usecase/command/user"
	userquery "skazitel-rus/internal/usecase/query/user"
	"skazitel-rus/pkg/response"
)

type UserHandler struct {
	registerUserHandler     *usercommand.RegisterUserHandler
	setUserOnlineHandler    *usercommand.SetUserOnlineHandler
	authenticateUserHandler *userquery.AuthenticateUserHandler
}

func NewUserHandler(
	registerUserHandler *usercommand.RegisterUserHandler,
	setUserOnlineHandler *usercommand.SetUserOnlineHandler,
	authenticateUserHandler *userquery.AuthenticateUserHandler,
) *UserHandler {
	return &UserHandler{
		registerUserHandler:     registerUserHandler,
		setUserOnlineHandler:    setUserOnlineHandler,
		authenticateUserHandler: authenticateUserHandler,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
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

	cmd := usercommand.RegisterUserCommand{
		Username: req.Username,
		Password: req.Password,
	}

	err = h.registerUserHandler.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Ошибка регистрации")
		return
	}

	response.SuccessResponse(w, http.StatusCreated, nil, "Пользователь зарегистрирован")
}

func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
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

	q := userquery.AuthenticateUserQuery{
		Username: req.Username,
		Password: req.Password,
	}

	isValid, err := h.authenticateUserHandler.Handle(r.Context(), q)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	if !isValid {
		response.ErrorResponse(w, http.StatusUnauthorized, "Неверные учетные данные")
		return
	}

	response.SuccessResponse(w, http.StatusOK, nil, "Аутентификация успешна")
}

func (h *UserHandler) SetOnline(w http.ResponseWriter, r *http.Request) {
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

	cmd := usercommand.SetUserOnlineCommand{
		Username: req.Username,
		IsOnline: req.IsOnline,
	}

	err = h.setUserOnlineHandler.Handle(r.Context(), cmd)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, nil, "Статус обновлен")
}
