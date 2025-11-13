package registerusecase

import (
	"context"
	"skazitel-rus/internal/domain/user"
)

type RegisterUserCommand struct {
	Username string
	Password string
}

type RegisterUserHandler struct {
	userRepo user.UserRepository
}

func NewRegisterUserHandler(userRepo user.UserRepository) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo: userRepo,
	}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, cmd RegisterUserCommand) error {
	return h.userRepo.RegisterUser(cmd.Username, cmd.Password)
}
