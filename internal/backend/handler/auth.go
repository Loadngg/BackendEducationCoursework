package handler

import (
	"net/http"

	"backend/internal/backend/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
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

func (h *Handler) SignIn(c *gin.Context) {
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
