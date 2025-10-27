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

func TestMetricUpdateHandler(t *testing.T) {
	testStorage := storage.NewMemStorage()
	r := router.SetupRoutes(testStorage)

	type body struct {
		ID    string `json:"id"`
		MType string `json:"type"`
		Value any    `json:"value"`
		Delta any    `json:"delta"`
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		body   body
		want   string
		status int
	}{
		{"should create gauge metric", body{ID: "TestGauge", MType: "gauge", Value: 1.5}, `{"metrics":{"id":"TestGauge","type":"gauge","value": 1.5},"success":true}`, 200},
		{"should update gauge metric", body{ID: "TestGauge", MType: "gauge", Value: 1.5}, `{"metrics":{"id":"TestGauge","type":"gauge","value": 1.5},"success":true}`, 200},
		{"should create counter metric", body{ID: "TestCounter", MType: "counter", Delta: 1}, `{"metrics":{"id":"TestCounter","type":"counter","delta":1},"success":true}`, 200},
		{"should update counter metric", body{ID: "TestCounter", MType: "counter", Delta: 1}, `{"metrics":{"id":"TestCounter","type":"counter","delta":2},"success":true}`, 200},
		{"should return 400 if wrong metric type was provided", body{ID: "TestCounter", MType: "wrongtype", Delta: 1}, `{"message":"Key: 'Metrics.MType' Error:Field validation for 'MType' failed on the 'oneof' tag","success":false}`, 400},
		{"should return 400 if no metric type was provided", body{ID: "", MType: "counter", Delta: 1}, `{"message":"Key: 'Metrics.ID' Error:Field validation for 'ID' failed on the 'required' tag","success":false}`, 400},
		{"should return 400 if wrong metric value for counter was provided", body{ID: "TestWrongCounterValue", MType: "counter", Delta: 1.2}, `{"message":"json: cannot unmarshal number 1.2 into Go struct field Metrics.delta of type int64","success":false}`, 400},
		{"should return 400 if wrong metric value for gauge was provided", body{ID: "TestWrongCounterValue", MType: "gauge", Delta: "abc"}, `{"message":"json: cannot unmarshal string into Go struct field Metrics.delta of type int64","success":false}`, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			body, _ := json.Marshal(test.body)
			req, _ := http.NewRequest("POST", "/update", strings.NewReader(string(body)))

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

	r := router.SetupRoutes(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want   string
		status int
	}{
		{"should get metrics", fmt.Sprintf(`<html><body><b>%s</b>: %s<br/></body></html>`, testMetrics.ID, testMetrics.GetValue()), 200},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, test.status, w.Code)
		assert.Equal(t, test.want, w.Body.String())
	}
}

func TestMetricGetValueHandler(t *testing.T) {
	testDelta := int64(2)
	testMetrics := model.Metrics{
		ID:    "TestID",
		MType: model.Counter,
		Delta: &testDelta,
	}
	testStorage := storage.NewMemStorage()
	testStorage.UpdateMetrics(testMetrics)
	testMetricResponse, _ := json.Marshal(testMetrics)

	r := router.SetupRoutes(testStorage)

	type body struct {
		ID    string `json:"id"`
		MType string `json:"type"`
	}

	tests := []struct {
		name string // description of this test case
		body
		want   string
		status int
	}{
		{"should get metric value", body{ID: testMetrics.ID, MType: testMetrics.MType}, string(testMetricResponse), 200},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			body, _ := json.Marshal(test.body)
			req, _ := http.NewRequest("POST", "/value", strings.NewReader(string(body)))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.status, w.Code)
			assert.JSONEq(t, test.want, w.Body.String())

		})
	}
}
