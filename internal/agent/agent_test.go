package agent_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func TestAgent_UpdateSingleMetric(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRoutes(testStorage)

	srv := httptest.NewServer(testRouter)
	defer srv.Close()

	type body struct {
		ID    string  `json:"id"`
		MType string  `json:"type"`
		Value float64 `json:"value"`
	}

	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed},
		{method: http.MethodPost, expectedCode: http.StatusOK},
	}
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			// делаем запрос с помощью библиотеки resty к адресу запущенного сервера,
			// который хранится в поле URL соответствующей структуры
			req := resty.New().R().
				SetMethod(tc.method).
				SetURL(srv.URL + "/update").SetBody(body{ID: "TestID", MType: "gauge", Value: 1}).
				SetResult(&responses.MetricUpdateResponse{})

			res, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tc.expectedCode, res.StatusCode(), "Response code didn't match expected")
		})
	}
}
