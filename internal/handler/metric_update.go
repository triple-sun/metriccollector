package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/requests"
	"github.com/triple-sun/metriccollector/internal/storage"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func UpdateMetricJSONHandler(storage storage.MemStorageRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body requests.UpdateMetricRequest

		err := ctx.BindJSON(&body)

		if err != nil {
			logger.Log.Err(err).Msg(err.Error())
			if err.Error() == "Key: 'Metrics.ID' Error:Field validation for 'ID' failed on the 'required' tag" {
				ctx.AbortWithError(http.StatusNotFound, err)
			} else {
				ctx.AbortWithError(http.StatusBadRequest, err)
			}
			return

		}

		logger.Log.Info().Msgf("Получен запрос обновления метрики: %v", body)

		updated := storage.UpdateMetrics(model.Metrics{ID: body.ID, MType: body.MType, Value: body.Value, Delta: body.Delta})

		ctx.JSON(http.StatusOK, &updated)

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}

func UpdateMetricHandler(storage storage.MemStorageRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		ctx.JSON(http.StatusOK, &updated)

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}
