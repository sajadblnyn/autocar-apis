package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{service: services.NewCountryService(cfg)}
}

func (h *CountryHandler) CreateCountry(c *gin.Context) {
	Create(c, h.service.CreateCountry)
}

func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	Update(c, h.service.UpdateCountry)
}

func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	Delete(c, h.service.DeleteCountry)
}

func (h *CountryHandler) GetCountry(c *gin.Context) {
	GetById(c, h.service.GetCountry)
}

func (h *CountryHandler) GetCountriesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCountryByFilter)
}
