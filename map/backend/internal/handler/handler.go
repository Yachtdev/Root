package handler

import (
	"encoding/json"
	"net/http"
	"server/gen"
)

type (
	// Handler реализует обработчик запросов.
	Handler struct {
		service Service
	}

	// Service реализует сервисный слой.
	Service interface {
		SendAisPointToMqtt(lat, lon, course float32)
	}
)

// nolint:grouper
var _ gen.ServerInterface = (*Handler)(nil)

// New handler-constructor.
func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) sendError(w http.ResponseWriter, code int, message string) {
	err := gen.Error{
		Code:    code,
		Message: message,
	}

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}
