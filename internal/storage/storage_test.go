package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/triple-sun/metriccollector/internal/model"
)

func TestMemStorage_UpdateMetrics(t *testing.T) {
	testDelta := int64(1)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		metrics models.Metrics
		want    models.Metrics
		wantErr bool
	}{
		{"should update counter metric", models.Metrics{ID: "TestMetric", MType: "counter", Delta: &testDelta}, models.Metrics{ID: "TestMetric", MType: "counter", Delta: &testDelta}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMemStorage()
			got := ms.UpdateMetrics(tt.metrics)

			assert.Equal(t, tt.want, got)
		})
	}
}
