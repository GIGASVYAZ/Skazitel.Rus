package userquery

import "skazitel-rus/internal/domain/user"

type UserRepository interface {
	GetByUsername(username string) (*user.User, error)
}
