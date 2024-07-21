package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{service: services.NewColorService(cfg)}
}

func (h *ColorHandler) CreateColor(c *gin.Context) {
	Create(c, h.service.CreateColor)
}

func (h *ColorHandler) UpdateColor(c *gin.Context) {
	Update(c, h.service.UpdateColor)
}

func (h *ColorHandler) DeleteColor(c *gin.Context) {
	Delete(c, h.service.DeleteColor)
}

func (h *ColorHandler) GetColor(c *gin.Context) {
	GetById(c, h.service.GetColor)
}

func (h *ColorHandler) GetColorsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetColorByFilter)
}
