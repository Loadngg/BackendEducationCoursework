package handler

import (
	"log/slog"

	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	UserIdentity(c *gin.Context)
}

type Authorization interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type Lectures interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetAllMaterials(c *gin.Context)
}

type Quiz interface {
}

type Results interface {
}

type Handler struct {
	Middleware
	Authorization
	Lectures
	Quiz
	Results
}

func New(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		Middleware:    NewMiddlewareHandler(service, log),
		Authorization: NewAuthorizationHandler(service.Authorization, log),
		Lectures:      NewLecturesHandler(service.Lectures, log),
	}
}
