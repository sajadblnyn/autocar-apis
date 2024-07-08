package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{userService: services.NewUSerService(cfg)}
}

func (u *UserHandler) SendOtp(c *gin.Context) {
	r := dto.GetOtpRequest{}

	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}

	err = u.userService.SendOtp(r)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateOtpErrorToStatusCode(err), helper.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 1))
}
