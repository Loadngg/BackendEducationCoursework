package server

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/handler"
	"backend/internal/config"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

const (
	envRelease = "prod"
	envDev     = "dev"
	envLocal   = "local"
)

type Server struct {
	gin *gin.Engine
}

func New(env string, log *slog.Logger, h *handler.Handler) *Server {
	s := &Server{}

	switch env {
	case envRelease:
		gin.SetMode(gin.ReleaseMode)
	case envDev, envLocal:
		gin.SetMode(gin.DebugMode)
	}
	s.gin = gin.New()

	s.gin.Use(sloggin.New(log))
	s.gin.Use(gin.Recovery())

	auth := s.gin.Group("/auth")
	{
		auth.POST("/sign-up", h.Authorization.SignUp)
		auth.POST("/sign-in", h.Authorization.SignIn)
	}

	api := s.gin.Group("/api", h.Middleware.UserIdentity)
	{
		lectures := api.Group("/lectures")
		{
			lectures.GET("/", h.Lectures.GetAll)
			lectures.GET("/:id", h.Lectures.GetById)
			lectures.GET("/:id/materials", h.Lectures.GetAllMaterials)
		}
		quizzes := api.Group("/quizzes")
		{
			quizzes.GET("/")
			quizzes.GET("/:id")
			questions := quizzes.Group("/:id/questions")
			{
				questions.GET("/:question_id")
				questions.POST("/:question_id/answers")
			}
		}
		api.GET("/results")
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
