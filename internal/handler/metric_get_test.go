package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func TestGetAllMetricsHandler(t *testing.T) {
	testDelta := int64(2)
	testMetrics := model.Metrics{
		ID:    "TestMetrics",
		MType: model.Counter,
		Delta: &testDelta,
	}
	testStorage := storage.NewMemStorage()
	testStorage.UpdateMetrics(testMetrics)

	testRouter := router.SetupRouter(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want   string
		status int
	}{
		{"should get metrics", fmt.Sprintf(`<html><body><b>%s</b>: %s<br/></body></html>`, testMetrics.ID, testMetrics.GetValue()), 200},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			testRouter.ServeHTTP(w, req)
			assert.Equal(t, test.status, w.Code)
			assert.Equal(t, test.want, w.Body.String())
		})

	}
}

func TestGetMetricValueJSONHandler(t *testing.T) {
	testDelta := int64(2)
	testMetrics := model.Metrics{
		ID:    "TestID",
		MType: model.Counter,
		Delta: &testDelta,
	}
	testStorage := storage.NewMemStorage()
	testStorage.UpdateMetrics(testMetrics)
	testMetricResponse, _ := json.Marshal(testMetrics)

	testRouter := router.SetupRouter(testStorage)

	tests := []struct {
		name   string // description of this test case
		body   string
		want   string
		status int
	}{
		{"should get metric value", `{"id": "TestID", "type": "counter"}`, string(testMetricResponse), 200},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/value", strings.NewReader(test.body))

			req.Header.Set("Content-Type", "application/json")

			testRouter.ServeHTTP(w, req)

			assert.Equal(t, test.status, w.Code)
			assert.JSONEq(t, test.want, w.Body.String())

		})
	}
}
