package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{service: services.NewCarTypeService(cfg)}
}

func (h *CarTypeHandler) CreateCarType(c *gin.Context) {
	Create(c, h.service.CreateCarType)
}

func (h *CarTypeHandler) UpdateCarType(c *gin.Context) {
	Update(c, h.service.UpdateCarType)
}

func (h *CarTypeHandler) DeleteCarType(c *gin.Context) {
	Delete(c, h.service.DeleteCarType)
}

func (h *CarTypeHandler) GetCarType(c *gin.Context) {
	GetById(c, h.service.GetCarType)
}

func (h *CarTypeHandler) GetCarTypesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarTypeByFilter)
}
