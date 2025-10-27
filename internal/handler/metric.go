package handler

import (
	"fmt"
	"log"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/storage"
)

type GetMetricParams struct {
	
}

func GetAllMetricsHandler(storage *storage.MemStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result string

		metrics := storage.GetMetrics()

		for name, metric := range *metrics {
			result += fmt.Sprintf("<b>%s</b>: %s<br/>", name, metric.GetValue())
		}

		ctx.Data(200, "text/html; charset=utf-8", fmt.Appendf(nil, `<html><body>%s</body></html>`, result))
	}
}

func MetricGetValueHandler(storage *storage.MemStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body model.Metrics

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		metric, findErr := storage.GetMetricsByID(body.ID)

		if findErr != nil {
			ctx.AbortWithError(http.StatusNotFound, findErr)
			return
		}

		ctx.JSON(http.StatusOK, metric)
	}
}

func MetricUpdateHandler(storage storage.MemStorageRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body model.Metrics

		logger.Log.Info().Interface("body", ctx.Request.GetBody)
		err := ctx.BindJSON(&body)

		if err != nil {
			if err.Error() == "Key: 'Metrics.ID' Error:Field validation for 'ID' failed on the 'required' tag" {
				ctx.AbortWithError(http.StatusNotFound, err)
			} else {
				ctx.AbortWithError(http.StatusBadRequest, err)
			}
			return
		}
		// обновляем
		updated := storage.UpdateMetrics(body)

		// назначаем ответ
		ctx.JSON(http.StatusOK, responses.MetricUpdateResponse{
			Success: true, Metrics: updated,
		})

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}
