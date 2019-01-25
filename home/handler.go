package home

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	now := time.Now().String()
	h.logger.Printf("request received: %v", now)
	writer.Write([]byte(now))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}