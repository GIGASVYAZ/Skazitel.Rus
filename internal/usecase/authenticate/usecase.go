package authenticateusecase

import (
	"context"
	"errors"
	"skazitel-rus/internal/domain/user"
	"skazitel-rus/internal/infrastructure/jwt"
)

type AuthenticateUserQuery struct {
	Username string
	Password string
}

type AuthenticateUserResult struct {
	Token    string `json:"token"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type AuthenticateUserHandler struct {
	userRepo user.UserRepository
}

func NewAuthenticateUserHandler(userRepo user.UserRepository) *AuthenticateUserHandler {
	return &AuthenticateUserHandler{
		userRepo: userRepo,
	}
}

func (h *AuthenticateUserHandler) Handle(ctx context.Context, q AuthenticateUserQuery) (*AuthenticateUserResult, error) {
	user, err := h.userRepo.GetByUsername(q.Username)
	if err != nil {
		return nil, errors.New("Непредвиденная ошибка:" + err.Error())
	}

	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	if user.Password != q.Password {
		return nil, errors.New("пароль неверный")
	}

	token, err := jwt.GenerateToken(int(user.ID), user.Username)
	if err != nil {
		return nil, errors.New("ошибка при генерации токена: " + err.Error())
	}

	return &AuthenticateUserResult{
		Token:    token,
		UserID:   int(user.ID),
		Username: user.Username,
	}, nil
}
