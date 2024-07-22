package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelPersianYearHandler struct {
	service *services.CarModelPersianYearService
}

func NewCarModelPersianYearHandler(cfg *config.Config) *CarModelPersianYearHandler {
	return &CarModelPersianYearHandler{service: services.NewCarModelPersianYearService(cfg)}
}

func (h *CarModelPersianYearHandler) CreateCarModelPersianYear(c *gin.Context) {
	Create(c, h.service.CreateCarModelPersianYear)
}

func (h *CarModelPersianYearHandler) UpdateCarModelPersianYear(c *gin.Context) {
	Update(c, h.service.UpdateCarModelPersianYear)
}

func (h *CarModelPersianYearHandler) DeleteCarModelPersianYear(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelPersianYear)
}

func (h *CarModelPersianYearHandler) GetCarModelPersianYear(c *gin.Context) {
	GetById(c, h.service.GetCarModelPersianYear)
}

func (h *CarModelPersianYearHandler) GetCarModelPersianYearsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelPersianYearByFilter)
}
