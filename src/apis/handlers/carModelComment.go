package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{service: services.NewCarModelCommentService(cfg)}
}

func (h *CarModelCommentHandler) CreateCarModelComment(c *gin.Context) {
	Create(c, h.service.CreateCarModelComment)
}

func (h *CarModelCommentHandler) UpdateCarModelComment(c *gin.Context) {
	Update(c, h.service.UpdateCarModelComment)
}

func (h *CarModelCommentHandler) DeleteCarModelComment(c *gin.Context) {
	Delete(c, h.service.DeleteCarModelComment)
}

func (h *CarModelCommentHandler) GetCarModelComment(c *gin.Context) {
	GetById(c, h.service.GetCarModelComment)
}

func (h *CarModelCommentHandler) GetCarModelCommentsByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCarModelCommentByFilter)
}
