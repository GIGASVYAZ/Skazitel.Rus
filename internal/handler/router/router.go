package router

import (
	"log"
	"net/http"

	jwt "skazitel-rus/internal/infrastructure/jwt"
	messagerepository "skazitel-rus/internal/infrastructure/repository/message"
	userrepository "skazitel-rus/internal/infrastructure/repository/user"
	authenticateusecase "skazitel-rus/internal/usecase/authenticate"
	getmessageusecase "skazitel-rus/internal/usecase/getmessage"
	messageshandlerusecase "skazitel-rus/internal/usecase/messageshandler"
	notfoundusecase "skazitel-rus/internal/usecase/notfound"
	registerusecase "skazitel-rus/internal/usecase/register"
	sendmessageusecase "skazitel-rus/internal/usecase/sendmessage"
	setonlineusecase "skazitel-rus/internal/usecase/setonline"
	"skazitel-rus/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool, cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := userrepository.New(pool)
	messageRepo := messagerepository.New(pool)

	registerHandler := registerusecase.NewRegisterUserHandler(userRepo)
	setUserOnlineHandler := setonlineusecase.NewSetUserOnlineHandler(userRepo)
	authenticateHandler := authenticateusecase.NewAuthenticateUserHandler(userRepo, cfg.Server.TokenTTL)
	sendMessageHandler := sendmessageusecase.NewSendMessageHandler(messageRepo)
	getMessagesHandler := getmessageusecase.NewGetMessagesHandler(messageRepo)

	messageHandler := messageshandlerusecase.NewMessageHandler(sendMessageHandler, getMessagesHandler)

	mux.HandleFunc("/", notfoundusecase.NotFoundHandler)
	mux.HandleFunc("/users/register", registerHandler.Register)
	mux.HandleFunc("/users/authenticate", authenticateHandler.Authenticate)
	mux.HandleFunc("/users/set-online", setUserOnlineHandler.SetOnline)

	mux.HandleFunc("/messages", jwt.Middleware(messageHandler.MessagesHandler))

	log.Println("Маршруты успешно инициализированы")
	return mux
}
