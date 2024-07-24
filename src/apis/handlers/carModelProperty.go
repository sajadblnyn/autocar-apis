package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelPropertyHandler struct {
	service *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{service: services.NewCarModelPropertyService(cfg)}
}

func (h *CarModelPropertyHandler) CreateCarModelProperty(c *gin.Context) {
	Create(c, h.service.CreateCarModelProperty)
}

func (h *CarModelPropertyHandler) UpdateCarModelProperty(c *gin.Context) {
	Update(c, h.service.UpdateCarModelProperty)
}

func (h *CarModelPropertyHandler) DeleteCarModelProperty(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelProperty)
}

func (h *CarModelPropertyHandler) GetCarModelProperty(c *gin.Context) {
	GetById(c, h.service.GetCarModelProperty)
}

func (h *CarModelPropertyHandler) GetCarModelPropertiesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelPropertyByFilter)
}
