package userquery

import (
	"context"
	"errors"
)

type AuthenticateUserQuery struct {
	Username string
	Password string
}

type AuthenticateUserHandler struct {
	userRepo UserRepository
}

func NewAuthenticateUserHandler(userRepo UserRepository) *AuthenticateUserHandler {
	return &AuthenticateUserHandler{
		userRepo: userRepo,
	}
}

func (h *AuthenticateUserHandler) Handle(ctx context.Context, q AuthenticateUserQuery) (bool, error) {
	user, err := h.userRepo.GetByUsername(q.Username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("пользователь не найден")
	}

	if user.Password == q.Password {
		return true, nil
	}

	return false, errors.New("пароль неверный")
}
