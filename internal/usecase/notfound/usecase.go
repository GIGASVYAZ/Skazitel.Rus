package notfoundusecase

import (
	"net/http"
	"skazitel-rus/pkg/response"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.ErrorResponse(w, http.StatusNotFound, "Маршрут не найден")
}
