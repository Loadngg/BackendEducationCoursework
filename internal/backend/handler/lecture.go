package handler

import (
	"log/slog"
	"net/http"

	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

type LecturesHandler struct {
	log     *slog.Logger
	service service.Lectures
}

func NewLecturesHandler(service service.Lectures, log *slog.Logger) *LecturesHandler {
	return &LecturesHandler{
		service: service,
		log:     log,
	}
}

func (h *LecturesHandler) GetAll(c *gin.Context) {
	lectures, err := h.service.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(lectures) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No lectures found"})
		return
	}

	c.JSON(http.StatusOK, lectures)
}

func (h *LecturesHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	lecture, err := h.service.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *LecturesHandler) GetAllMaterials(c *gin.Context) {
	lectureId := c.Param("id")

	materials, err := h.service.GetAllMaterials(lectureId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	if len(materials) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No materials found"})
		return
	}

	c.JSON(http.StatusOK, materials)
}
