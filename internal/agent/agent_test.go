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

func TestAgent_GetUpdateMetricRequest(t *testing.T) {
	testStorage := storage.NewMemStorage()
	testRouter := router.SetupRouter(testStorage)

	srv := httptest.NewServer(testRouter)
	defer srv.Close()

	testParams := agent.MetricUpdateParams{ID: "TestMetric", MType: "counter", Delta: 1}

	type want struct {
		path            string
		contentType     string
		contentEncoding string
	}

	tests := []struct {
		name   string
		method string
		params agent.MetricUpdateParams
		want   want
	}{
		{name: "should create update metrics request with compression", params: testParams, want: want{path: "/update", contentType: "application/json", contentEncoding: "gzip"}}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// делаем запрос с помощью библиотеки resty к адресу запущенного сервера,
			// который хранится в поле URL соответствующей структуры
			req, _ := agent.GetUpdateMetricRequest(srv.URL, test.params)

			assert.Equal(t, test.want.path, req.URL.RequestURI(), "Request URL didn't match expected")
			assert.Equal(t, test.want.contentType, req.Header.Get("Content-Type"), "Content-Type header didn't match expected")
			assert.Equal(t, test.want.contentEncoding, req.Header.Get("Content-Encoding"))
		})
	}
}

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
			req, _ := agent.GetUpdateMetricRequest(srv.URL, test.params)
			req.Header.Set("Accept-Encoding", "gzip")

			res, err := testClient.Do(req)
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
