package setonlineusecase

import (
	"context"
	"skazitel-rus/internal/domain/user"
)

type SetUserOnlineCommand struct {
	Username string
	IsOnline bool
}

type SetUserOnlineHandler struct {
	userRepo user.UserRepository
}

func NewSetUserOnlineHandler(userRepo user.UserRepository) *SetUserOnlineHandler {
	return &SetUserOnlineHandler{
		userRepo: userRepo,
	}
}

func (h *SetUserOnlineHandler) Handle(ctx context.Context, cmd SetUserOnlineCommand) error {
	return h.userRepo.UpdateIsOnline(cmd.Username, cmd.IsOnline)
}
