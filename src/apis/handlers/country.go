package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
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
	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}

	res, err := h.service.CreateCountry(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := h.service.UpdateCountry(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := h.service.DeleteCountry(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, int(helper.Success)))
}

func (h *CountryHandler) GetCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := h.service.GetCountry(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

func (h *CountryHandler) GetCountriesByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}

	res, err := h.service.GetCountryByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}
