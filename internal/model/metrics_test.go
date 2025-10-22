package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/model"
)

func TestMetrics_GetValue(t *testing.T) {
	var testDelta int64 = 123
	var testValue float64 = 12.3

	tests := []struct {
		name   string // description of this test case
		metric model.MetricsModel
		want   string
	}{
		{name: "should get Counter metric value", metric: model.Metrics{ID: "test", MType: model.Counter, Delta: &testDelta}, want: "123"},
		{name: "should get Gauge metric value", metric: model.Metrics{ID: "test", MType: model.Gauge, Value: &testValue}, want: "12.3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.metric.GetValue())
		})
	}
}
