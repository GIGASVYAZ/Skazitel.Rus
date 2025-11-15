package authenticateusecase

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain/user"
)

type AuthenticateUserQuery struct {
	Username string
	Password string
}

type AuthenticateUserHandler struct {
	userRepo user.UserRepository
}

func NewAuthenticateUserHandler(userRepo user.UserRepository) *AuthenticateUserHandler {
	return &AuthenticateUserHandler{
		userRepo: userRepo,
	}
}

func (h *AuthenticateUserHandler) Handle(ctx context.Context, q AuthenticateUserQuery) (bool, error) {
	user, err := h.userRepo.GetByUsername(q.Username)
	if err != nil {
		return false, errors.New("Непредвиденная ошибка:" + err.Error())
	}

	if user == nil {
		return false, errors.New("пользователь не найден")
	}

	if user.Password == q.Password {
		return true, nil
	}

	return false, errors.New("пароль неверный")
}
