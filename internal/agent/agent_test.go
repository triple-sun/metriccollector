package agent_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/agent"
	models "github.com/triple-sun/metriccollector/internal/model"
)

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestAgent_UpdateSingleMetric(t *testing.T) {
	testMType := models.Gauge
	testMName := "TestMetric"
	testMValue := "1"

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://localhost:8080/update")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: io.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	testAgent := agent.NewAgent(client, "https://localhost:8080")
	body, err := testAgent.UpdateSingleMetric(testMType, testMName, testMValue)
	assert.NoError(t, err)
	assert.Equal(t, []byte("OK"), body)
}
