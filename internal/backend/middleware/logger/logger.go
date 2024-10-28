package logger

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func New(log *slog.Logger) gin.HandlerFunc {
	return sloggin.New(log)
}
