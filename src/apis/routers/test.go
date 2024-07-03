package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
	"github.com/sajadblnyn/autocar-apis/apis/middlewares"
)

func Test(r *gin.RouterGroup) {
	TestHandler := handlers.NewTestHandler()

	r.GET("/query", TestHandler.BindQuery)
	r.GET("/query-array", TestHandler.BindQueryArray)

	r.POST("/header1", TestHandler.BindHeader1)
	r.POST("/header2", TestHandler.BindHeader2)

	r.POST("/json-body", TestHandler.BindJsonBody)
	r.POST("/form-body", TestHandler.BindForm)
	r.POST("/file", TestHandler.BindFormFile)

	r.GET("/uri/:id", TestHandler.BindUri)

	r.GET("/middleware", middlewares.TestMiddleware(), TestHandler.BindQuery)

}
