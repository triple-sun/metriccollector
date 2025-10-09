package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/triple-sun/metriccollector/internal/handler"
	storage "github.com/triple-sun/metriccollector/internal/storage"
)

func TestMetricHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	type path struct {
		mtype  string
		mname  string
		mvalue string
	}

	testStorage := storage.NewMemStorage()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		storage storage.IMemStorage
		path    path
		want    want
	}{
		{"should create gauge metric", testStorage, path{"gauge", "TestGauge", "1"}, want{200, `{"id": "TestGauge", "type": "gauge", "value": 1}`, "application/json"}},
		{"should update gauge metric", testStorage, path{"gauge", "TestGauge", "1"}, want{200, `{"id": "TestGauge", "type": "gauge", "value": 1}`, "application/json"}},
		{"should create counter metric", testStorage, path{"counter", "TestCounter", "1"}, want{200, `{"id": "TestCounter", "type": "counter", "delta": 1}`, "application/json"}},
		{"should update counter metric", testStorage, path{"counter", "TestCounter", "1"}, want{200, `{"id": "TestCounter", "type": "counter", "delta": 2}`, "application/json"}},
		{"should return 400 if wrong metric type was found", testStorage, path{"wrongtype", "TestCounter", "1"}, want{400, "Некорректный тип метрики: wrongtype\n", "text/plain; charset=utf-8"}},
		{"should return 400 if no metric name type was found", testStorage, path{"counter", "", "1"}, want{404, "Не найдено имя метрики\n", "text/plain; charset=utf-8"}},
		{"should return 400 if wrong metric value for counter was found", testStorage, path{"counter", "TestCounter", "1.5"}, want{400, "Ошибка разбора значения метрики: некорректный тип значения для типа Counter: strconv.ParseInt: parsing \"1.5\": invalid syntax\n", "text/plain; charset=utf-8"}},
		{"should return 400 if wrong metric value for gauge was found", testStorage, path{"gauge", "TestGauge", "abc"}, want{400, "Ошибка разбора значения метрики: некорректный тип значения для типа Gauge: strconv.ParseFloat: parsing \"abc\": invalid syntax\n", "text/plain; charset=utf-8"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/update", nil)

			request.SetPathValue("mtype", test.path.mtype)
			request.SetPathValue("mname", test.path.mname)
			request.SetPathValue("mvalue", test.path.mvalue)

			// создаём новый Recorder
			w := httptest.NewRecorder()
			handlerFunc := handler.MetricHandler(test.storage)

			handlerFunc(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)

			if json.Valid(resBody) {
				assert.JSONEq(t, test.want.response, string(resBody))
			} else {
				assert.Equal(t, test.want.response, string(resBody))
			}

			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
