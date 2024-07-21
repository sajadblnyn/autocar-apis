package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{service: services.NewCarModelColorService(cfg)}
}

func (h *CarModelColorHandler) CreateCarModelColor(c *gin.Context) {
	Create(c, h.service.CreateCarModelColor)
}

func (h *CarModelColorHandler) UpdateCarModelColor(c *gin.Context) {
	Update(c, h.service.UpdateCarModelColor)
}

func (h *CarModelColorHandler) DeleteCarModelColor(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelColor)
}

func (h *CarModelColorHandler) GetCarModelColor(c *gin.Context) {
	GetById(c, h.service.GetCarModelColor)
}

func (h *CarModelColorHandler) GetCarModelColorsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelColorByFilter)
}
