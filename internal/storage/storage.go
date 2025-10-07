package server

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	models "github.com/triple-sun/metriccollector/internal/model"
)

type MemStorage struct {
	metrics []models.Metrics
}

type IMemStorage interface {
	UpdateCounter(name string, value string) error
	UpdateGauge(name string, value string) error
	GetMetrics() []models.Metrics
	GetMetricString() string
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make([]models.Metrics, 0),
	}
}

func (m *MemStorage) UpdateCounter(name string, value string) error {
	parsed, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Counter: %s", err.Error())
		return errors.New(message)
	}

	index := slices.IndexFunc(m.metrics, func(metric models.Metrics) bool {
		return metric.ID == name && metric.MType == models.Counter
	})

	if index == -1 {
		newDelta := parsed
		m.metrics = append(m.metrics, models.Metrics{ID: name, MType: models.Counter, Delta: &newDelta})
	} else {
		fmt.Printf("Найдена метрика %s: [%d]\n", name, index)

		newDelta := parsed + *m.metrics[index].Delta
		m.metrics[index].Delta = &newDelta
	}

	return nil
}

func (m *MemStorage) UpdateGauge(name string, value string) error {
	parsed, err := strconv.ParseFloat(value, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Gauge: %s", err.Error())
		return errors.New(message)
	}

	index := slices.IndexFunc(m.metrics, func(metric models.Metrics) bool {
		return metric.ID == name && metric.MType == models.Gauge
	})

	if index == -1 {
		m.metrics = append(m.metrics, models.Metrics{ID: name, MType: models.Gauge, Value: &parsed})
	} else {
		fmt.Printf("Найдена метрика %s: [%d]\n", name, index)

		m.metrics[index].Value = &parsed
	}

	return nil

}

func (m *MemStorage) GetMetrics() []models.Metrics {
	return m.metrics
}

func (m *MemStorage) GetMetricString() string {
	return fmt.Sprintf("Metrics: %s", fmt.Sprint(m.metrics))
}
