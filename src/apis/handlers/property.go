package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{service: services.NewPropertyService(cfg)}
}

func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	Create(c, h.service.CreateProperty)
}

func (h *PropertyHandler) UpdateProperty(c *gin.Context) {
	Update(c, h.service.UpdateProperty)
}

func (h *PropertyHandler) DeleteProperty(c *gin.Context) {
	Delete(c, h.service.DeleteProperty)
}

func (h *PropertyHandler) GetProperty(c *gin.Context) {
	GetById(c, h.service.GetProperty)
}

func (h *PropertyHandler) GetPropertiesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetPropertyByFilter)
}
