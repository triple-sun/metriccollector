package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type ServerConfig struct {
	Address string `env:"ADDRESS"`
}

func ParseServerFlags(address *string) {
	flag.StringVar(address, "a", "localhost:8080", "Адрес в формате host:port")
	flag.Parse()
}

func ParseServerEnv(config *ServerConfig, address *string) {
	envErr := env.Parse(config)
	if envErr != nil {
		log.Fatal(envErr)
	}

	if config.Address != "" {
		*address = config.Address
	}
}
