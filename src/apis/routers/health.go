package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
)

func Health(r *gin.RouterGroup) {

	healthHandler := handlers.NewHealthHandler()
	r.GET("/", healthHandler.Health)
}
