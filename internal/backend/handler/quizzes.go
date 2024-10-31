package handler

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/models"
	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

type QuizzesHandler struct {
	log     *slog.Logger
	service service.Quizzes
}

func NewQuizzesHandler(service service.Quizzes, log *slog.Logger) *QuizzesHandler {
	return &QuizzesHandler{
		service: service,
		log:     log,
	}
}

func (h *QuizzesHandler) GetAll(c *gin.Context) {
	quizzes, err := h.service.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(quizzes) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No quizzes found"})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

func (h *QuizzesHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	lecture, err := h.service.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *QuizzesHandler) GetAllQuestions(c *gin.Context) {
	id := c.Param("id")

	questions, err := h.service.GetAllQuestions(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(questions) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No questions found"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *QuizzesHandler) GetQuestionById(c *gin.Context) {
	id := c.Param("question_id")

	lecture, err := h.service.GetQuestionById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *QuizzesHandler) GetAllAnswers(c *gin.Context) {
	questionId := c.Param("question_id")

	answers, err := h.service.GetAllAnswers(questionId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(answers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No answers found"})
		return
	}

	c.JSON(http.StatusOK, answers)
}

func (h *QuizzesHandler) CreateAnswer(c *gin.Context) {
	var input models.UserAnswers

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	*input.UserID, _ = GetUserId(c)

	id, err := h.service.CreateAnswer(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
