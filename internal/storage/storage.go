package storage

import (
	"fmt"

	models "github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/requests"
)

type MemStorage struct {
	metrics map[string]models.Metrics
}

type MemStorageRepository interface {
	UpdateMetrics(metrics models.Metrics) models.Metrics
	GetMetrics() *map[string]models.Metrics
	FindOne(params *requests.GetMetricValueRequest) (*models.Metrics, error)
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

func (ms *MemStorage) GetMetrics() *map[string]models.Metrics {
	return &ms.metrics
}

func (ms *MemStorage) FindOne(params *requests.GetMetricValueRequest) (*models.Metrics, error) {
	found, ok := ms.metrics[params.ID]

	if !ok || found.MType != params.MType {
		return nil, fmt.Errorf("метрика %s не найдена", params.ID)
	}

	return &found, nil
}
