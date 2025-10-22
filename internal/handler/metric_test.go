package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func TestMetricUpdateHandler(t *testing.T) {
	testStorage := storage.NewMemStorage()
	r := router.Setup(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		url    string
		method string
		want   string
		status int
	}{
		{"should create gauge metric", "/update/gauge/TestGauge/1", "POST", `{"metrics":{"id":"TestGauge","type":"gauge","value": 1},"success":true}`, 200},
		{"should update gauge metric", "/update/gauge/TestGauge/1", "POST", `{"metrics":{"id":"TestGauge","type":"gauge","value": 1},"success":true}`, 200},
		{"should create counter metric", "/update/counter/TestCounter/1", "POST", `{"metrics":{"id":"TestCounter","type":"counter","delta":1},"success":true}`, 200},
		{"should update counter metric", "/update/counter/TestCounter/1", "POST", `{"metrics":{"id":"TestCounter","type":"counter","delta":2},"success":true}`, 200},
		{"should return 400 if wrong metric type was provided", "/update/wrongtype/TestWrongType/1", "POST", `{"message":"некорректный тип метрики: wrongtype","success":false}`, 400},
		{"should return 400 if no metric type was provided", "/update/counter//1", "POST", `{"message":"не найдено имя метрики","success":false}`, 404},
		{"should return 400 if wrong metric value for counter was provided", "/update/counter/TestCounter/1.5", "POST", `{"message":"некорректный тип значения","success":false}`, 400},
		{"should return 400 if wrong metric value for gauge was provided", "/update/gauge/TestGauge/abc", "POST", `{"message":"некорректный тип значения","success":false}`, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(test.method, test.url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, test.status, w.Code)
			assert.JSONEq(t, test.want, w.Body.String())
		})
	}
}

func TestGetAllMetricsHandler(t *testing.T) {
	testDelta := int64(2)
	testMetrics := model.Metrics{
		ID:    "TestMetrics",
		MType: model.Counter,
		Delta: &testDelta,
	}
	testStorage := storage.NewMemStorage()
	testStorage.UpdateMetrics(testMetrics)

	r := router.Setup(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		url    string
		method string
		want   string
		status int
	}{
		{"should get metrics", "/", "GET", fmt.Sprintf(`<html><body><b>%s</b>: %s<br/></body></html>`, testMetrics.ID, testMetrics.GetValue()), 200},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.method, test.url, nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, test.status, w.Code)
		assert.Equal(t, test.want, w.Body.String())
	}
}
