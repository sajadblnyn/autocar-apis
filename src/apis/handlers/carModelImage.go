package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{service: services.NewCarModelImageService(cfg)}
}

func (h *CarModelImageHandler) CreateCarModelImage(c *gin.Context) {
	Create(c, h.service.CreateCarModelImage)
}

func (h *CarModelImageHandler) UpdateCarModelImage(c *gin.Context) {
	Update(c, h.service.UpdateCarModelImage)
}

func (h *CarModelImageHandler) DeleteCarModelImage(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelImage)
}

func (h *CarModelImageHandler) GetCarModelImage(c *gin.Context) {
	GetById(c, h.service.GetCarModelImage)
}

func (h *CarModelImageHandler) GetCarModelImagesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelImageByFilter)
}
