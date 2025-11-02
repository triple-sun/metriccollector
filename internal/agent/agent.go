package agent

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/triple-sun/metriccollector/internal/model"
)

type MetricUpdateParams struct {
	ID    string
	MType string
	Value float64
	Delta int64
}

func SendUpdateMetricRequest(client *http.Client, address string, params MetricUpdateParams) (*http.Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(model.Metrics{ID: params.ID, MType: params.MType, Delta: &params.Delta, Value: &params.Value})

	if err != nil {
		return nil, err
	}
	zw := gzip.NewWriter(&buf)
	defer zw.Close()
	if _, err := zw.Write(body); err != nil {
		return nil, fmt.Errorf("error compressing data: %v", err)
	}
	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("failed to compress data: %v", err)
	}
	req, err := http.NewRequest("POST", address+"/update", &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Accept-Encoding", "gzip")

	return client.Do(req)
}
