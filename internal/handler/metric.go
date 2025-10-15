package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/storage"
	"github.com/triple-sun/metriccollector/internal/utils"
)

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

func MetricUpdateHandler(storage storage.MemStorageRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var response responses.MetricUpdateResponse

		mtype := ctx.Param("mtype")
		mname := ctx.Param("mname")
		mvalue := ctx.Param("mvalue")

		// Проверяем тип метрик
		err := utils.ValidateMType(mtype)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Проверяем имя метрики
		err = utils.ValidateMName(mname)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		log.Printf("Получен запрос обновления метрики: Type: [%s]; ID: [%s]; Value: [%s]\n", mtype, mname, mvalue)

		metrics := model.Metrics{ID: mname, MType: mtype}

		// Обновляем метрики
		switch mtype {
		case model.Gauge:
			parsed, err := strconv.ParseFloat(mvalue, 64)
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("некорректный тип значения"))
				return
			}
			metrics.Value = &parsed
		case model.Counter:
			parsed, err := strconv.ParseInt(mvalue, 10, 64)
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("некорректный тип значения"))
				return
			}
			metrics.Delta = &parsed
		}

		// обновляем
		updated := storage.UpdateMetrics(metrics)

		// назначаем ответ
		ctx.JSON(http.StatusOK, responses.MetricUpdateResponse{
			Success: true, Metrics: updated,
		})

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}

func MetricGetValueHandler(storage *storage.MemStorage) gin.HandlerFunc {
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

		metric, findErr := storage.GetMetricsByName(mname)

		if findErr != nil {
			ctx.AbortWithError(http.StatusNotFound, findErr)
		}

		ctx.Data(200, "text/plain; charset=utf-8", []byte(metric.GetValue()))
	}
}
