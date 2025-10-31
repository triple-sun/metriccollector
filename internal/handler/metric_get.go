package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/requests"
	"github.com/triple-sun/metriccollector/internal/storage"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func GetAllMetricsHandler(storage *storage.MemStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result string

		fmt.Println("Got all metrics")

		metrics := storage.GetMetrics()

		for name, metric := range *metrics {
			result += fmt.Sprintf("<b>%s</b>: %s<br/>", name, metric.GetValue())
		}

		ctx.Data(200, "text/html; charset=utf-8", fmt.Appendf(nil, `<html><body>%s</body></html>`, result))
	}
}


func GetMetricValueHandler(storage *storage.MemStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mtype := ctx.Param("mtype")
		mname := ctx.Param("mname")

		typeErr := utils.ValidateMType(mtype)
		nameErr := utils.ValidateMName(mname)

		if typeErr != nil {
			ctx.AbortWithError(http.StatusBadRequest, typeErr)
			return
		}

		if nameErr != nil {
			ctx.AbortWithError(http.StatusBadRequest, typeErr)
			return
		}

		metric, findErr := storage.FindOne(&requests.GetMetricValueRequest{ID: mname, MType: mtype})

		if findErr != nil {
			ctx.AbortWithError(http.StatusNotFound, findErr)
			return
		}

		ctx.JSON(http.StatusOK, &metric)
	}
}

func GetMetricValueJSONHandler(storage *storage.MemStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.GetMetricValueRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		metric, err := storage.FindOne(&req)

		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		ctx.JSON(http.StatusOK, &metric)
	}
}
