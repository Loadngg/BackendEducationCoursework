package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"backend/internal/backend/service"
	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

type MiddlewareHandler struct {
	log     *slog.Logger
	service *service.Service
}

func NewMiddlewareHandler(service *service.Service, log *slog.Logger) *MiddlewareHandler {
	return &MiddlewareHandler{
		service: service,
		log:     log,
	}
}

func (h *MiddlewareHandler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "missing authorization header", h.log)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization header", h.log)
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error(), h.log)
		return
	}

	c.Set(userCtx, userId)
}

func GetUserId(c *gin.Context) (int32, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int32)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
