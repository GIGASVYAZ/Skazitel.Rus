package messagecommand

type MessageRepository interface {
	Create(userID int64, content string) error
}
