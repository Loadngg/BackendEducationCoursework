package handler

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

type ResultsHandler struct {
	log     *slog.Logger
	service service.Results
}

func NewResultsHandler(service service.Results, log *slog.Logger) *ResultsHandler {
	return &ResultsHandler{
		service: service,
		log:     log,
	}
}

func (h *ResultsHandler) GetAll(c *gin.Context) {
	results, err := h.service.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No results found"})
		return
	}

	c.JSON(http.StatusOK, results)
}
