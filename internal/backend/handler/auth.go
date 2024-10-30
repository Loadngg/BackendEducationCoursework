package handler

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/models"
	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

type AuthorizationHandler struct {
	log     *slog.Logger
	service service.Authorization
}

func NewAuthorizationHandler(service service.Authorization, log *slog.Logger) *AuthorizationHandler {
	return &AuthorizationHandler{
		service: service,
		log:     log,
	}
}

func (h *AuthorizationHandler) SignUp(c *gin.Context) {
	var input models.Users

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthorizationHandler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	token, err := h.service.GenerateToken(input.Login, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
