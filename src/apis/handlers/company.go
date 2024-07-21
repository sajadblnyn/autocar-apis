package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{service: services.NewCompanyService(cfg)}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	Create(c, h.service.CreateCompany)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	Update(c, h.service.UpdateCompany)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	Delete(c, h.service.DeleteCompany)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	GetById(c, h.service.GetCompany)
}

func (h *CompanyHandler) GetCompaniesByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetCompanyByFilter)
}
