package main

import (
	"context"
	"log"
	app "skazitel-rus/internal/handler"
	"skazitel-rus/pkg/database"
)

const DATABASE_URL = "postgres://postgres:mysecretpassword@localhost:5432/postgres"

func main() {
	ctx := context.Background()

	err := database.InitPool(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Ошибка инициализации пула:", err)
	}
	defer database.ClosePool()

	app.RunServer()
}
