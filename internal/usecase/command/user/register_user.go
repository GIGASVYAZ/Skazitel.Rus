package usercommand

import "context"

type RegisterUserCommand struct {
	Username string
	Password string
}

type RegisterUserHandler struct {
	userRepo UserRepository
}

func NewRegisterUserHandler(userRepo UserRepository) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo: userRepo,
	}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, cmd RegisterUserCommand) error {
	return h.userRepo.Create(cmd.Username, cmd.Password)
}
