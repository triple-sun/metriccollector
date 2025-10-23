package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		logger.Log.
			Info().
			Str("method", ctx.Request.Method).
			Str("url", ctx.Request.URL.RequestURI()).
			Dur("ms", time.Since(start)).
			Msg("incoming request")
	}
}
