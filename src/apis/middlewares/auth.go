package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
	"github.com/sajadblnyn/autocar-apis/services"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := services.NewTokenService(cfg)
	var err error
	var claimMap map[string]interface{}
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if len(token) < 2 || auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[0])

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.ExpiredToke}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.InvalidToken}

				}
			}
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}

		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		ctx.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		ctx.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		ctx.Set(constants.EmailKey, claimMap[constants.EmailKey])
		ctx.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])
		ctx.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		ctx.Next()

	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if len(ctx.Keys) < 1 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, 300))
			return
		}

		var exists bool
		rolesVal, exists := ctx.Keys[constants.RolesKey]

		if !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, 301))
			return
		}
		roles := rolesVal.([]interface{})

		mapRoles := make(map[string]int)
		for _, v := range roles {
			mapRoles[v.(string)] = 0
		}

		for _, v := range validRoles {
			_, exists = mapRoles[v]
			if exists {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, 302))

	}
}
