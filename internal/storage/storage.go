package server

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"

	models "github.com/triple-sun/metriccollector/internal/model"
)

type MemStorage struct {
	metrics []models.Metrics
}

type IMemStorage interface {
	UpdateCounter(name string, value string) (models.Metrics, error)
	UpdateGauge(name string, value string) (models.Metrics, error)
	GetMetrics() []models.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make([]models.Metrics, 0),
	}
}

func (m *MemStorage) UpdateCounter(mname string, mvalue string) (models.Metrics, error) {
	parsed, err := strconv.ParseInt(mvalue, 10, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Counter: %s", err.Error())
		return models.Metrics{}, errors.New(message)
	}

	index := slices.IndexFunc(m.metrics, func(metric models.Metrics) bool {
		return metric.ID == mname && metric.MType == models.Counter
	})

	if index == -1 {
		newDelta := parsed
		newMetric := models.Metrics{ID: mname, MType: models.Counter, Delta: &newDelta}
		m.metrics = append(m.metrics, newMetric)

		log.Printf("Добавлена метрика %s\n", mname)

		return newMetric, nil
	} else {
		log.Printf("Найдена метрика %s: [%d]\n", mname, index)

		newDelta := parsed + *m.metrics[index].Delta
		m.metrics[index].Delta = &newDelta
		return m.metrics[index], nil

	}

}

func (m *MemStorage) UpdateGauge(mname string, mvalue string) (models.Metrics, error) {
	parsed, err := strconv.ParseFloat(mvalue, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Gauge: %s", err.Error())
		return models.Metrics{}, errors.New(message)
	}

	index := slices.IndexFunc(m.metrics, func(metric models.Metrics) bool {
		return metric.ID == mname && metric.MType == models.Gauge
	})

	if index == -1 {
		newMetric := models.Metrics{ID: mname, MType: models.Gauge, Value: &parsed}
		m.metrics = append(m.metrics, newMetric)

		log.Printf("Добавлена метрика %s\n", mname)

		return newMetric, nil

	} else {
		log.Printf("Найдена метрика %s: [%d]\n", mname, index)

		m.metrics[index].Value = &parsed

		return m.metrics[index], nil
	}
}

func (m *MemStorage) GetMetrics() []models.Metrics {
	return m.metrics
}
