package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/triple-sun/metriccollector/internal/model"
	storage "github.com/triple-sun/metriccollector/internal/storage"
)

func MetricHandler(storage storage.IMemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var metric models.Metrics

		mtype := r.PathValue("mtype")
		mname := r.PathValue("mname")
		mvalue := r.PathValue("mvalue")

		if mtype != models.Counter && mtype != models.Gauge {
			http.Error(w, fmt.Sprintf(`Некорректный тип метрики: %s`, mtype), http.StatusBadRequest)
			return
		}

		if mname == "" {
			http.Error(w, `Не найдено имя метрики`, http.StatusNotFound)
			return
		}

		log.Printf("Получен запрос обновления метрики: Type: [%s]; ID: [%s]; Value: [%s]\n", mtype, mname, mvalue)

		switch mtype {
		case models.Gauge:
			updated, err := storage.UpdateGauge(mname, mvalue)
			if err != nil {
				http.Error(w, fmt.Sprintf(`Ошибка разбора значения метрики: %s`, err.Error()), http.StatusBadRequest)
				return
			}

			/** Назначаем ответ */
			metric = updated
		case models.Counter:
			updated, err := storage.UpdateCounter(mname, mvalue)
			if err != nil {
				http.Error(w, fmt.Sprintf(`Ошибка разбора значения метрики: %s`, err.Error()), http.StatusBadRequest)
				return
			}
			/** Назначаем ответ */
			metric = updated
		}

		w.Header().Set("Content-Type", "application/json")

		resJSON, encErr := json.Marshal(metric)

		if encErr != nil {
			http.Error(w, "ошибка преобразования в json", http.StatusInternalServerError)
		}

		_, writeErr := w.Write(resJSON)

		if writeErr != nil {
			http.Error(w, "ошибка записи ответа", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)

		log.Printf("Запрос обновления метрики обработан успешно!\n")
	}
}
