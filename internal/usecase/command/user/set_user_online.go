package usercommand

import "context"

type SetUserOnlineCommand struct {
	Username string
	IsOnline bool
}

type SetUserOnlineHandler struct {
	userRepo UserRepository
}

func NewSetUserOnlineHandler(userRepo UserRepository) *SetUserOnlineHandler {
	return &SetUserOnlineHandler{
		userRepo: userRepo,
	}
}

func (h *SetUserOnlineHandler) Handle(ctx context.Context, cmd SetUserOnlineCommand) error {
	return h.userRepo.UpdateIsOnline(cmd.Username, cmd.IsOnline)
}
