package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/helper"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/limiter"

	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, (helper.OtpLimiterError), errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
