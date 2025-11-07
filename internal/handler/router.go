package handler

import (
	"log"
	"net/http"
	"skazitel-rus/internal/repository/messageRepository"
	"skazitel-rus/internal/repository/userRepository"
	"skazitel-rus/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := userRepository.NewPostgresUserRepository()
	messageRepo := messageRepository.NewPostgresMessageRepository()

	userUseCase := usecase.NewUserUseCase(userRepo)
	messageUseCase := usecase.NewMessageUseCase(messageRepo)

	handler := NewHandler(messageUseCase, userUseCase)

	mux.HandleFunc("/", handler.NotFoundHandler)
	mux.HandleFunc("/messages", handler.MessagesHandler)

	log.Println("Маршруты успешно инициализированы")

	return mux
}
