package handler

import (
	"log/slog"

	"backend/internal/backend/service"
)

type Handler struct {
	log     *slog.Logger
	service *service.Service
}

func New(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}
