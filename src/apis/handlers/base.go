package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

var logger logging.Logger

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, (helper.ValidationError), err))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, (helper.InternalError), err))
		return

	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, (helper.Success)))

}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := new(Ti)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, (helper.ValidationError), err))
		return
	}
	res, err := caller(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, (helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, (helper.Success)))
}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, (helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, (helper.Success)))
}

func GetById[To any](c *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, (helper.InternalError), err))
		return

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, (helper.Success)))
}

func GetByFilter[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, (helper.ValidationError), err))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, (helper.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, (helper.Success)))
}
