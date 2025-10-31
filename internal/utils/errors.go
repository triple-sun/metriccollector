package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/responses"
)

func RespondWithError(ctx *gin.Context) {
	// Step3: Use the last error
	err := ctx.Errors.Last().Err
	// Step4: Respond with a generic error message
	ctx.JSON(ctx.Writer.Status(), responses.ErrorResponse{
		Success: false,
		Message: err.Error(),
	})
}
