package handler

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
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

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
