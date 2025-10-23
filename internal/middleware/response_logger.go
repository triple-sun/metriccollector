package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
)

func ResponseLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()

		logger.Log.
			Info().
			Int("status", ctx.Writer.Status()).
			Int("size", ctx.Writer.Size()).
			Msg("response")
	}
}
