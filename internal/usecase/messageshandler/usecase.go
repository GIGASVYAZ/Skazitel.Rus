package messageshandlerusecase

import (
	"net/http"
	getmessageusecase "skazitel-rus/internal/usecase/getmessage"
	sendmessageusecase "skazitel-rus/internal/usecase/sendmessage"
	"skazitel-rus/pkg/response"
)

type MessageHandler struct {
	sendMessageHandler *sendmessageusecase.SendMessageHandler
	getMessagesHandler *getmessageusecase.GetMessagesHandler
}

func NewMessageHandler(
	sendMessageHandler *sendmessageusecase.SendMessageHandler,
	getMessagesHandler *getmessageusecase.GetMessagesHandler,
) *MessageHandler {
	return &MessageHandler{
		sendMessageHandler: sendMessageHandler,
		getMessagesHandler: getMessagesHandler,
	}
}

func (h *MessageHandler) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.SendMessage(w, r)
	case http.MethodGet:
		h.GetMessages(w, r)
	default:
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Метод не разрешен")
	}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	h.sendMessageHandler.SendMessage(w, r)
}

func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	h.getMessagesHandler.GetMessages(w, r)
}
