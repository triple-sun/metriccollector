package storage

import (
	"fmt"
	"log"

	models "github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/utils"
)

type MemStorage struct {
	metrics map[string]models.Metrics
}

type MemStorageRepository interface {
	UpdateCounter(name string, value string) (models.Metrics, error)
	UpdateGauge(name string, value string) (models.Metrics, error)
	UpdateMetrics(metrics models.Metrics) models.Metrics
	GetMetrics() *map[string]models.Metrics
	GetMetricsByName(mname string) (models.Metrics, error)
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]models.Metrics, 0),
	}
}

func (ms *MemStorage) UpdateMetrics(metrics models.Metrics) models.Metrics {
	found, ok := ms.metrics[metrics.ID]

	if !ok {
		ms.metrics[metrics.ID] = metrics
	} else {
		if found.Delta != nil {
			oldDelta := *found.Delta
			newDelta := oldDelta + *metrics.Delta
			metrics.Delta = &newDelta
		}

		ms.metrics[metrics.ID] = metrics
	}

	return ms.metrics[metrics.ID]
}

func (ms *MemStorage) UpdateCounter(mname string, mvalue string) (models.Metrics, error) {
	parsed, err := utils.ParseCounterValue(mvalue)

	if err != nil {
		return models.Metrics{}, err
	}

	found, ok := ms.metrics[mname]

	if !ok {
		newDelta := parsed
		ms.metrics[mname] = models.Metrics{ID: mname, MType: models.Counter, Delta: &newDelta}

		log.Printf("Добавлена метрика %s\n", mname)

		return ms.metrics[mname], nil
	} else {
		log.Printf("Найдена метрика %s \n", mname)

		newDelta := parsed + *found.Delta
		found.Delta = &newDelta

		log.Printf("Обновлена метрика %s \n", mname)

		return found, nil
	}
}

func (ms *MemStorage) UpdateGauge(mname string, mvalue string) (models.Metrics, error) {
	parsed, err := utils.ParseGaugeValue(mvalue)

	if err != nil {
		return models.Metrics{}, err
	}

	found, ok := ms.metrics[mname]

	if !ok {
		ms.metrics[mname] = models.Metrics{ID: mname, MType: models.Gauge, Value: &parsed}

		log.Printf("Добавлена метрика %s\n", mname)

		return ms.metrics[mname], nil

	} else {
		log.Printf("Найдена метрика %s\n", mname)

		found.Value = &parsed

		return found, nil
	}
}

func (ms *MemStorage) GetMetrics() *map[string]models.Metrics {
	return &ms.metrics
}

func (ms *MemStorage) GetMetricsByName(mname string) (models.Metrics, error) {
	found, ok := ms.metrics[mname]

	if !ok {
		return models.Metrics{}, fmt.Errorf("метрика %s не найдена", mname)
	}

	return found, nil
}
