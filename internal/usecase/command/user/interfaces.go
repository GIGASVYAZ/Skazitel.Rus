package usercommand

type UserRepository interface {
	Create(username string, password string) error
	UpdateIsOnline(username string, isOnline bool) error
}
