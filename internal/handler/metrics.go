package handler

import (
	"fmt"
	"net/http"

	models "github.com/triple-sun/metriccollector/internal/model"
	storage "github.com/triple-sun/metriccollector/internal/storage"
)

func HandleMetric(w http.ResponseWriter, r *http.Request, storage storage.IMemStorage) {
	if r.Method != http.MethodPost {
		fmt.Println(`Метод отличается от POST`)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	metricType := r.PathValue("metricType")
	metricName := r.PathValue("metricName")
	metricValue := r.PathValue("metricValue")

	if metricName == "" {
		fmt.Printf(`Не найдено имя метрики`)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Printf("Получен запрос обновления метрики: %s, %s, %s\n", metricName, metricType, metricValue)

	switch metricType {
	case models.Gauge:
		err := storage.UpdateGauge(metricName, metricValue)
		if err != nil {
			fmt.Printf(`Ошибка разбора значения метрики: %s\n`, err.Error())

			w.WriteHeader(http.StatusBadRequest)

			return
		}
	case models.Counter:
		err := storage.UpdateCounter(metricName, metricValue)
		if err != nil {
			fmt.Printf(`Ошибка разбора значения метрики: %s`, err.Error())

			w.WriteHeader(http.StatusBadRequest)

			return
		}
	default:
		fmt.Printf(`Некорректный тип метрики: %s`, metricType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Метрики обновлены: %s \n", storage.GetMetricString())

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// установим правильный заголовок для типа данных
	// пока установим ответ-заглушку, без проверки ошибок
	_, _ = w.Write([]byte(storage.GetMetricString()))

	fmt.Printf("Запрос обновления метрики обработан успешно!")

}
