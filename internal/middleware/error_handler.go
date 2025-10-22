package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/responses"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(ctx.Errors) > 0 {
			// Step3: Use the last error
			err := ctx.Errors.Last().Err

			// Step4: Respond with a generic error message
			ctx.JSON(ctx.Writer.Status(), responses.ErrorResponse{
				Success:   false,
				Message:   err.Error(),
			})
		}
	}
}
