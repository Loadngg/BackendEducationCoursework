package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, log *slog.Logger) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
