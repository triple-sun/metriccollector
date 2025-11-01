package handler_test

import (
	"bytes"
	"compress/gzip"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func TestUpdateMetricJSONHandler(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRouter(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		body   string
		want   string
		status int
	}{
		{"should create gauge metric", `{"id": "TestGauge", "type": "gauge", "value": 1.5}`, `{"id":"TestGauge","type":"gauge","value": 1.5}`, 200},
		{"should update gauge metric", `{"id": "TestGauge", "type": "gauge", "value": 1.5}`, `{"id":"TestGauge","type":"gauge","value": 1.5}`, 200},
		{"should create counter metric", `{"id": "TestCounter", "type": "counter", "Delta": 1}`, `{"id":"TestCounter","type":"counter","delta":1}`, 200},
		{"should update counter metric", `{"id": "TestCounter", "type": "counter", "Delta": 1}`, `{"id":"TestCounter","type":"counter","delta":2}`, 200},
		{"should return 400 if wrong metric type was provided", `{"id": "TestCounter", "type": "wrongtype", "Delta": 1}`, `{"message":"некорректный тип метрики: wrongtype"}`, 400},
		{"should return 404 if no metric name was provided", `{"id": "", "type": "counter", "Delta": 1}`, `{"message":"не найдено имя метрики"}`, 404},
		{"should return 400 if wrong metric value for counter was provided", `{"id": "TestWrongCounterValue", "type": "counter", "Delta": 1.2}`, `{"message":"json: cannot unmarshal number 1.2 into Go struct field UpdateMetricRequest.delta of type int64"}`, 400},
		{"should return 400 if wrong metric value for gauge was provided", `{"id": "TestWrongCounterValue", "type": "gauge", "Delta": "abc"}`, `{"message":"json: cannot unmarshal string into Go struct field UpdateMetricRequest.delta of type int64"}`, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			zb := gzip.NewWriter(buf)
			_, err := zb.Write([]byte(test.body))
			require.NoError(t, err)
			err = zb.Close()
			require.NoError(t, err)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/update", buf)
			req.Header.Set("Content-Encoding", "gzip")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept-Encoding", "")
			testRouter.ServeHTTP(w, req)

			assert.Equal(t, test.status, w.Code)
			assert.JSONEq(t, test.want, w.Body.String())
		})
	}
}

func TestUpdateMetricHandler(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRouter(testStorage)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		url    string
		method string
		want   string
		status int
	}{
		{"should create gauge metric", "/update/gauge/TestGauge/1", "POST", `{"id":"TestGauge","type":"gauge","value": 1}`, 200},
		{"should update gauge metric", "/update/gauge/TestGauge/1", "POST", `{"id":"TestGauge","type":"gauge","value": 1}`, 200},
		{"should create counter metric", "/update/counter/TestCounter/1", "POST", `{"id":"TestCounter","type":"counter","delta":1}`, 200},
		{"should update counter metric", "/update/counter/TestCounter/1", "POST", `{"id":"TestCounter","type":"counter","delta":2}`, 200},
		{"should return 400 if wrong metric type was provided", "/update/wrongtype/TestWrongType/1", "POST", `{"message":"некорректный тип метрики: wrongtype"}`, 400},
		{"should return 404 if no metric name was provided", "/update/counter//1", "POST", `{"message":"не найдено имя метрики"}`, 404},
		{"should return 400 if wrong metric value for counter was provided", "/update/counter/TestCounter/1.5", "POST", `{"message": "некорректный тип значения"}`, 400},
		{"should return 400 if wrong metric value for gauge was provided", "/update/gauge/TestGauge/abc", "POST", `{"message": "некорректный тип значения"}`, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, test.url, nil)
			
			testRouter.ServeHTTP(w, req)

			assert.Equal(t, test.status, w.Code)
			assert.JSONEq(t, test.want, w.Body.String())
		})
	}
}
