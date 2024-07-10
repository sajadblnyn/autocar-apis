package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

func RecoveryErrors(c *gin.Context, err any) {
	logger := logging.NewLogger(config.GetConfig())
	er, ok := err.(error)
	if ok {
		logger.Error(logging.General, logging.RecoverError, fmt.Sprintf("%v", err), nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), er))
		return
	}

	logger.Error(logging.General, logging.RecoverError, fmt.Sprintf("%v", err), nil)
	c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithAnyError(nil, false, int(helper.InternalError), err))

}
