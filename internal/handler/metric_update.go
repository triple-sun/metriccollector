package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/requests"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func UpdateMetricJSONHandler(storage storage.MemStorageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.UpdateMetricRequest{}

		if err := render.Bind(r, req); err != nil {
			if err.Error() == "не найдено имя метрики" {
				_ = render.Render(w, r, responses.NewErrorResponse(404, err))
			} else {
				_ = render.Render(w, r, responses.NewErrorResponse(400, err))
			}
			return
		}

		logger.Log.Info().Msgf("Получен запрос обновления метрики: %v", req)

		updated := storage.UpdateMetrics(model.Metrics{ID: req.ID, MType: req.MType, Value: req.Value, Delta: req.Delta})

		render.JSON(w, r, &updated)

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}

func UpdateMetricHandler(storage storage.MemStorageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.UpdateMetricRequest{ID: chi.URLParam(r, "mname"), MType: chi.URLParam(r, "mtype")}

		if err := req.ParseValue(chi.URLParam(r, "mvalue")); err != nil {
			_ = render.Render(w, r, responses.NewErrorResponse(400, err))
			return
		}

		if err := req.Validate(); err != nil {
			if err.Error() == "не найдено имя метрики" {
				_ = render.Render(w, r, responses.NewErrorResponse(404, err))
			} else {
				_ = render.Render(w, r, responses.NewErrorResponse(400, err))
			}
			return
		}

		log.Printf("Получен запрос обновления метрики: %v", req)

		// обновляем
		updated := storage.UpdateMetrics(model.Metrics{ID: req.ID, MType: req.MType, Value: req.Value, Delta: req.Delta})

		// назначаем ответ
		render.JSON(w, r, &updated)

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}
