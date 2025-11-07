package usecase

import (
	"skazitel-rus/internal/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) RegisterUser(username string, password string) error {
	return uc.userRepo.Create(username, password)
}

func (uc *UserUseCase) AuthenticateUser(username string, password string) (bool, error) {
	return uc.userRepo.IsPasswordByUsernameEqualTo(username, password)
}

func (uc *UserUseCase) SetUserOnline(username string, isOnline bool) error {
	return uc.userRepo.UpdateIsOnline(username, isOnline)
}
