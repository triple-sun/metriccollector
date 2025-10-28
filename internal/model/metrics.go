package model

import (
	"strconv"
)

const (
	Counter = "counter"
	Gauge   = "gauge"
)

// NOTE: Не усложняем пример, вводя иерархическую вложенность структур.
// Органичиваясь плоской моделью.
// Delta и Value объявлены через указатели,
// что бы отличать значение "0", от не заданного значения
// и соответственно не кодировать в структуру.
type Metrics struct {
	ID    string   `json:"id" binding:"required"`
	MType string   `json:"type" binding:"required,oneof=counter gauge"`
	Delta *int64   `json:"delta,omitempty" binding:"omitempty,required_if=MType counter"`
	Value *float64 `json:"value,omitempty" binding:"omitempty,required_if=MType gauge"`
	Hash  string   `json:"hash,omitempty"`
}

type MetricsModel interface {
	GetValue() string
}

func (m Metrics) GetValue() string {
	switch m.MType {
	case Counter:
		return strconv.FormatInt(*m.Delta, 10)
	case Gauge:
		return strconv.FormatFloat(*m.Value, 'f', -1, 64)
	default:
		return ""
	}
}
