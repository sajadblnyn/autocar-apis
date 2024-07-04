package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sajadblnyn/autocar-apis/apis/middlewares"
	"github.com/sajadblnyn/autocar-apis/apis/routers"
	"github.com/sajadblnyn/autocar-apis/apis/validations"
	"github.com/sajadblnyn/autocar-apis/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("iran-mobile-validator", validations.IranianMobileValidator, true)
		v.RegisterValidation("password", validations.PasswordValidator, true)

	}
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery() /*, middlewares.TestMiddleware()*/)

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test := v1.Group("/test")
		// test.Use(middlewares.TestMiddleware())
		routers.Test(test)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}
