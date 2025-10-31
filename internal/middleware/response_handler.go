package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func ResponseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		isJSON := strings.Contains(ctx.GetHeader("Content-Type"), "application/json")
		isHTML := strings.Contains(ctx.GetHeader("Content-Type"), "text/html")
		supportsGzip := strings.Contains(ctx.GetHeader("Accept-Encoding"), "gzip")

		ctx.Next()

		logger.Log.
			Info().
			Int("status", ctx.Writer.Status()).
			Int("size", ctx.Writer.Size()).
			Msg("response")
		// Step2: Check if any errors were added to the context
		if len(ctx.Errors) > 0 {
			logger.Log.Error().Msg(ctx.Errors.Last().Error())
			// Step3: Use the last error
			err := ctx.Errors.Last().Err
			// Step4: Respond with a generic error message
			ctx.JSON(ctx.Writer.Status(), &responses.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		} else if (isJSON || isHTML) && supportsGzip {
			logger.Log.Info().Msg("Compressing response body...")
			// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
			cw := utils.NewCompressWriter(ctx.Writer)
			// меняем оригинальный http.ResponseWriter на новый
			ctx.Header("Content-Encoding", "gzip")
			ctx.Writer = cw
			// не забываем отправить клиенту все сжатые данные после завершения middleware
			defer func() {
				err := cw.Close()
				if err != nil {
					logger.Log.Error().Err(err)
				}
				cw.Header().Add("Content-Encoding", "gzip")
			
				logger.Log.Info().Msg("Response body compressed!")
			}()
		}
	}
}
