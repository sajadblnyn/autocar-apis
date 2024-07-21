package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{service: services.NewGearboxService(cfg)}
}

func (h *GearboxHandler) CreateGearbox(c *gin.Context) {
	Create(c, h.service.CreateGearbox)
}

func (h *GearboxHandler) UpdateGearbox(c *gin.Context) {
	Update(c, h.service.UpdateGearbox)
}

func (h *GearboxHandler) DeleteGearbox(c *gin.Context) {
	Delete(c, h.service.DeleteGearbox)
}

func (h *GearboxHandler) GetGearbox(c *gin.Context) {
	GetById(c, h.service.GetGearbox)
}

func (h *GearboxHandler) GetGearboxesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetGearboxByFilter)
}
