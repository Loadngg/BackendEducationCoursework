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

type Quizzes interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetAllQuestions(c *gin.Context)
	GetQuestionById(c *gin.Context)
	GetAllAnswers(c *gin.Context)
	CreateAnswer(c *gin.Context)
}

type Results interface {
	GetAll(c *gin.Context)
}

type Handler struct {
	Middleware
	Authorization
	Lectures
	Quizzes
	Results
}

func New(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		Middleware:    NewMiddlewareHandler(service, log),
		Authorization: NewAuthorizationHandler(service.Authorization, log),
		Lectures:      NewLecturesHandler(service.Lectures, log),
		Quizzes:       NewQuizzesHandler(service.Quizzes, log),
		Results:       NewResultsHandler(service.Results, log),
	}
}
