package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{service: services.NewCarModelService(cfg)}
}

func (h *CarModelHandler) CreateCarModel(c *gin.Context) {
	Create(c, h.service.CreateCarModel)
}

func (h *CarModelHandler) UpdateCarModel(c *gin.Context) {
	Update(c, h.service.UpdateCarModel)
}

func (h *CarModelHandler) DeleteCarModel(c *gin.Context) {
	Delete(c, h.service.DeleteCarModel)
}

func (h *CarModelHandler) GetCarModel(c *gin.Context) {
	GetById(c, h.service.GetCarModel)
}

func (h *CarModelHandler) GetCarModelsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelByFilter)
}
