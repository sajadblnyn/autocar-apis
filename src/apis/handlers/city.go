package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{service: services.NewCityService(cfg)}
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	Create(c, h.service.CreateCity)
}

func (h *CityHandler) UpdateCity(c *gin.Context) {
	Update(c, h.service.UpdateCity)
}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	Delete(c, h.service.DeleteCity)
}

func (h *CityHandler) GetCity(c *gin.Context) {
	GetById(c, h.service.GetCity)
}

func (h *CityHandler) GetCitiesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCityByFilter)
}
