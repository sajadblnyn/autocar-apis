package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/pkg/metrics"
)

func Prometheus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		start := time.Now()
		path := ctx.FullPath()
		method := ctx.Request.Method
		ctx.Next()

		status := ctx.Writer.Status()
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).Observe(float64(time.Since(start) / time.Millisecond))
	}
}
