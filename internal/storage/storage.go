package server

import (
	"errors"
	"fmt"
	"strconv"
)

type MemStorage struct {
	counters map[string]int64
	gauges   map[string]float64
}

type IMemStorage interface {
	UpdateCounter(name string, value string) error
	UpdateGauge(name string, value string) error
	GetCounters() map[string]int64
	GetGauges() map[string]float64
	GetMetricString() string
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		counters: make(map[string]int64),
		gauges:   make(map[string]float64),
	}
}

func (m *MemStorage) UpdateCounter(name string, value string) error {
	converted, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Gauge: %s", err.Error())
		return errors.New(message)
	}

	m.counters[name] += converted

	return nil
}

func (m *MemStorage) UpdateGauge(name string, value string) error {
	val, err := strconv.ParseFloat(value, 64)

	if err != nil {
		message := fmt.Sprintf("некорректный тип значения для типа Counter: %s", err.Error())
		return errors.New(message)
	}

	m.gauges[name] = val

	return nil

}

func (m *MemStorage) GetCounters() map[string]int64 {
	return m.counters
}

func (m *MemStorage) GetGauges() map[string]float64 {
	return m.gauges
}

func (m *MemStorage) GetMetricString() string {
	return fmt.Sprintf("Counters: %s, Gauges: %s", fmt.Sprint(m.counters), fmt.Sprint(m.gauges))
}
