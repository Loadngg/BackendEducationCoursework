package server

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/handler"
	"backend/internal/backend/middleware/logger"
	"backend/internal/config"
	"github.com/gin-gonic/gin"
)

const (
	envRelease = "prod"
	envDev     = "dev"
	envLocal   = "local"
)

type Server struct {
	gin *gin.Engine
}

func New(env string, log *slog.Logger, handler *handler.Handler) *Server {
	s := &Server{}

	switch env {
	case envRelease:
		gin.SetMode(gin.ReleaseMode)
	case envDev, envLocal:
		gin.SetMode(gin.DebugMode)
	}
	s.gin = gin.New()

	s.gin.Use(logger.New(log))
	s.gin.Use(gin.Recovery())

	api := s.gin.Group("/api")
	{
		api.GET("/hello", handler.Hello)
	}

	return s
}

func (s *Server) Run(cfg *config.Config) error {
	httpServer := &http.Server{
		Addr:           cfg.Address,
		Handler:        s.gin,
		IdleTimeout:    cfg.IdleTimeout,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		MaxHeaderBytes: 1 << 20,
	}

	return httpServer.ListenAndServe()
}
