package router

import (
	"log"
	"net/http"

	messagerepository "skazitel-rus/internal/infrastructure/repository/message"
	userrepository "skazitel-rus/internal/infrastructure/repository/user"
	"skazitel-rus/internal/interfaces/handler"
	messagecommand "skazitel-rus/internal/usecase/command/message"
	usercommand "skazitel-rus/internal/usecase/command/user"
	messagequery "skazitel-rus/internal/usecase/query/message"
	userquery "skazitel-rus/internal/usecase/query/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := userrepository.New()
	messageRepo := messagerepository.New()

	registerHandler := usercommand.NewRegisterUserHandler(userRepo)
	setUserOnlineHandler := usercommand.NewSetUserOnlineHandler(userRepo)

	authenticateHandler := userquery.NewAuthenticateUserHandler(userRepo)

	sendMessageHandler := messagecommand.NewSendMessageHandler(messageRepo)

	getMessagesHandler := messagequery.NewGetMessagesHandler(messageRepo)

	userHandler := handler.NewUserHandler(registerHandler, setUserOnlineHandler, authenticateHandler)
	messageHandler := handler.NewMessageHandler(sendMessageHandler, getMessagesHandler)

	mux.HandleFunc("/", handler.NotFoundHandler)
	mux.HandleFunc("/users/register", userHandler.Register)
	mux.HandleFunc("/users/authenticate", userHandler.Authenticate)
	mux.HandleFunc("/users/set-online", userHandler.SetOnline)
	mux.HandleFunc("/messages", messageHandler.MessagesHandler)

	log.Println("Маршруты успешно инициализированы")
	return mux
}
