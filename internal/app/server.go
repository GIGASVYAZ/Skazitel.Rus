package app

import (
	"encoding/json"
	"log"
	"net/http"
	"skazitel-rus/internal/database/pg/chat"
	"strconv"
)

type SendMessageRequest struct {
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
}

type MessageResponse struct {
	Status string `json:"status"`
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sendMessageHandler(w, r)
	} else if r.Method == http.MethodGet {
		getMessagesHandler(w, r)
	} else {
		notFoundHandler(w, r)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "СLOX"})
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(MessageResponse{"Неверный метод"})
		return
	}

	var req SendMessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(MessageResponse{"Неверный формат данных"})
		return
	}

	err = chat.CreateMessage(req.UserId, req.Content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MessageResponse{"Ошибка при создании сообщения"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MessageResponse{"Сообщение отправлено"})
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(MessageResponse{"Метод не разрешен"})
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(MessageResponse{"Неверный параметр limit"})
			return
		}
	}

	messages, err := chat.GetMessages(limit)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MessageResponse{"Ошибка при получении сообщений"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func RunServer() {

	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/messages", messagesHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
