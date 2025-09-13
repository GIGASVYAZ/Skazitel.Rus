package app

import (
	"fmt"
	"skazitel-rus/internal/database/pg/chat"
)

func RunServer() {
	fmt.Println("не работает")
	userAuthenticate, err := chat.UserAuthenticate("LoL123", "Kek")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userAuthenticate)

	userAuthenticate, err = chat.UserAuthenticate("LoL", "KeK")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userAuthenticate)
	fmt.Println("работает")

	err = chat.UpdateUserStatus("LoL", false)
	if err != nil {
		fmt.Println(err)
	}
}
