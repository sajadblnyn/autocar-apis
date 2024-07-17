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

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{service: services.NewCityService(cfg)}
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}

	res, err := h.service.CreateCity(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func (h *CityHandler) UpdateCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := h.service.UpdateCity(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := h.service.DeleteCity(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, int(helper.Success)))
}

func (h *CityHandler) GetCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := h.service.GetCity(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

func (h *CityHandler) GetCitiesByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, int(helper.ValidationError), err))
		return
	}

	res, err := h.service.GetCityByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}
