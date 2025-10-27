package config_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/triple-sun/metriccollector/internal/config"
)

func TestParseAgentEnv(t *testing.T) {
	var cfg config.AgentConfig

	address := "http://testhost:8080"
	pInterval := 2
	rInterval := 10

	tests := []struct {
		name      string // description of this test case
		address   string
		pInterval int
		rInterval int
	}{
		{name: "should some env params", address: "", pInterval: 10, rInterval: 100},
		{name: "should set all env params", address: "localhost:8080", pInterval: 10, rInterval: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ADDRESS", tt.address)
			t.Setenv("POLL_INTERVAL", strconv.Itoa(tt.pInterval))
			t.Setenv("REPORT_INTERVAL", strconv.Itoa(tt.rInterval))

			config.ParseAgentEnv(&cfg, &address, &pInterval, &rInterval)

			assert.Equal(t, cfg.Address, tt.address)
			assert.Equal(t, cfg.PollInterval, tt.pInterval)
			assert.Equal(t, cfg.ReportInterval, tt.rInterval)

		})
	}
}
