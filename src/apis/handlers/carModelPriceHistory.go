package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelPriceHistoryHandler struct {
	service *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{service: services.NewCarModelPriceHistoryService(cfg)}
}

func (h *CarModelPriceHistoryHandler) CreateCarModelPriceHistory(c *gin.Context) {
	Create(c, h.service.CreateCarModelPriceHistory)
}

func (h *CarModelPriceHistoryHandler) UpdateCarModelPriceHistory(c *gin.Context) {
	Update(c, h.service.UpdateCarModelPriceHistory)
}

func (h *CarModelPriceHistoryHandler) DeleteCarModelPriceHistory(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelPriceHistory)
}

func (h *CarModelPriceHistoryHandler) GetCarModelPriceHistory(c *gin.Context) {
	GetById(c, h.service.GetCarModelPriceHistory)
}

func (h *CarModelPriceHistoryHandler) GetCarModelPriceHistoriesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelPriceHistoryByFilter)
}
