package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type AgentConfig struct {
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

func ParseAgentFlags(address *string, pInterval *int, rInterval *int) {
	flag.StringVar(address, "a", "localhost:8080", "Адрес в формате host:port")
	flag.IntVar(pInterval, "p", 2, "Интервал сбора метрик в секундах")
	flag.IntVar(rInterval, "r", 10, "Интервал отправки метрик в секундах")
	flag.Parse()
}

func ParseAgentEnv(config *AgentConfig, address *string, pInterval *int, rInterval *int) {
	envErr := env.Parse(config)
	
	if envErr != nil {
		log.Fatal(envErr)
	}

	if config.Address != "" {
		*address = config.Address
	}
	if config.PollInterval != 0 {
		*pInterval = config.PollInterval
	}
	if config.ReportInterval != 0 {
		*rInterval = config.ReportInterval
	}
}
