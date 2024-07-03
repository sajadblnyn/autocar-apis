package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func TestMiddleware(ctx *gin.Context) {
// 	token := ctx.GetHeader("token")
// 	if token != "1" {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"error": "permission denied",
// 		})
// 		return
// 	}
// 	ctx.Next()

// 	return
// }

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "1" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "permission denied",
			})
			return
		}
		ctx.Next()

		return
	}
}
