package main

import (
	"context"
	"log"
	"net/http"

	"skazitel-rus/internal/interfaces/router"
	"skazitel-rus/pkg/database"
	"skazitel-rus/pkg/httpserver"
)

func main() {
	cfg := httpserver.New()
	ctx := context.Background()

	err := database.InitPoolWithConfig(
		ctx,
		cfg.Database.URL,
		cfg.Database.MaxConns,
		cfg.Database.MinConns,
	)
	if err != nil {
		log.Fatal("Ошибка инициализации пула:", err)
	}

	defer database.ClosePool()

	pool := database.GetPool()
	if pool == nil {
		log.Fatal("Пул подключений не инициализирован")
	}

	mux := router.New(pool)
	runServer(mux, cfg.Server.Port)
}

func runServer(mux *http.ServeMux, port string) {
	log.Printf("Сервер запущен на http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
