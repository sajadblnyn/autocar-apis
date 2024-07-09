package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
	"github.com/sajadblnyn/autocar-apis/apis/middlewares"
	"github.com/sajadblnyn/autocar-apis/config"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	userHandler := handlers.NewUserHandler(cfg)

	r.POST("/send-otp", middlewares.OtpLimiter(cfg), userHandler.SendOtp)
}
