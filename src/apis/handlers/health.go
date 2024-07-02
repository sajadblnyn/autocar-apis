package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "im fine")
	return
}

func NewHealthHandler() (h *HealthHandler) {
	return &HealthHandler{}
}
