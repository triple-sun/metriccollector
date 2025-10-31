package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func RequestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("reqHandler")

		logger.Log.
			Info().
			Str("method", ctx.Request.Method).
			Str("url", ctx.Request.URL.RequestURI()).
			Dur("ms", time.Since(time.Now())).
			Msg("incoming request")

		isJSON := strings.Contains(ctx.GetHeader("Content-Type"), "application/json")
		isHTML := strings.Contains(ctx.GetHeader("Content-Type"), "text/html")
		supportsGzip := strings.Contains(ctx.GetHeader("Content-Encoding"), "gzip")

		if (isHTML || isJSON) && supportsGzip {
			logger.Log.Info().Msg("Decompressing request body...")

			cr, err := utils.NewCompressReader(ctx.Request.Body)

			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				ctx.Next()
				return
			}

			defer cr.Close()

			ctx.Request.Body = cr
		}

		ctx.Next()
	}
}
