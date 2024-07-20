package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{service: services.NewPropertyCategoryService(cfg)}
}

func (h *PropertyCategoryHandler) CreatePropertyCategory(c *gin.Context) {
	Create(c, h.service.CreatePropertyCategory)
}

func (h *PropertyCategoryHandler) UpdatePropertyCategory(c *gin.Context) {
	Update(c, h.service.UpdatePropertyCategory)
}

func (h *PropertyCategoryHandler) DeletePropertyCategory(c *gin.Context) {
	Delete(c, h.service.DeletePropertyCategory)
}

func (h *PropertyCategoryHandler) GetPropertyCategory(c *gin.Context) {
	GetById(c, h.service.GetPropertyCategory)
}

func (h *PropertyCategoryHandler) GetPropertyCategoriesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetPropertyCategoryByFilter)
}
