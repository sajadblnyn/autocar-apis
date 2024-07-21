package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
)

type HealthHandler struct {
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("im fine", true, 1))

}

func NewHealthHandler() (h *HealthHandler) {
	return &HealthHandler{}
}
