package agent_test

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/triple-sun/metriccollector/internal/agent"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)


func TestAgent_SendUpdateMetricRequest(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRouter(testStorage)

	srv := httptest.NewServer(testRouter)
	defer srv.Close()

	testClient := http.DefaultClient
	testParams := agent.MetricUpdateParams{ID: "TestMetric", MType: "counter", Delta: 1}

	type want struct {
		url             string
		contentType     string
		contentEncoding string
	}

	tests := []struct {
		name   string
		method string
		params agent.MetricUpdateParams
		want   want
	}{
		{name: "should send update metrics request with compression", params: testParams, want: want{url: srv.URL + "/update", contentType: "application/json", contentEncoding: "gzip"}}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// делаем запрос с помощью библиотеки resty к адресу запущенного сервера,
			// который хранится в поле URL соответствующей структуры
			res, err := agent.SendUpdateMetricRequest(testClient, srv.URL, test.params)
			require.NoError(t, err)

			defer res.Body.Close()

			fmt.Printf("%v", res.Header)

			zr, err := gzip.NewReader(res.Body)
			require.NoError(t, err)

			b, err := io.ReadAll(zr)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, res.StatusCode)
			fmt.Printf("%v\n", b)
		})
	}
}
