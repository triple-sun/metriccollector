package agent_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/agent"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func TestAgent_GetUpdateMetricRequest(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRoutes(testStorage)

	srv := httptest.NewServer(testRouter)
	defer srv.Close()

	testClient := resty.New().SetBaseURL(srv.URL)
	testParams :=  agent.MetricUpdateParams{ID: "TestMetric", MType: "counter", Delta: 1}

	tests := []struct {
		name         string
		method       string
		params       agent.MetricUpdateParams
		expectedBody string
		wantErr      bool
	}{
		{name: "should update metric", params: testParams, expectedBody: `{"id":"TestMetric","type":"counter","delta":1}`},
		{name: "should throw if no metric", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			// делаем запрос с помощью библиотеки resty к адресу запущенного сервера,
			// который хранится в поле URL соответствующей структуры
			req, err := agent.GetUpdateMetricRequest(testClient, test.params)

			if test.wantErr {
				assert.NoError(t, err, "error making HTTP request")
			} else {
				assert.JSONEq(t, test.expectedBody, string(req.Body.([]uint8)), "Request body didn't match expected")
			}
		})
	}
}
