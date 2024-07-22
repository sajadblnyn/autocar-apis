package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type PersianYearHandler struct {
	service *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{service: services.NewPersianYearService(cfg)}
}

func (h *PersianYearHandler) CreatePersianYear(c *gin.Context) {
	Create(c, h.service.CreatePersianYear)
}

func (h *PersianYearHandler) UpdatePersianYear(c *gin.Context) {
	Update(c, h.service.UpdatePersianYear)
}

func (h *PersianYearHandler) DeletePersianYear(c *gin.Context) {
	Delete(c, h.service.DeletePersianYear)
}

func (h *PersianYearHandler) GetPersianYear(c *gin.Context) {
	GetById(c, h.service.GetPersianYear)
}

func (h *PersianYearHandler) GetPersianYearsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetPersianYearByFilter)
}
