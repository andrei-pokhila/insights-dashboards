package config

import (
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

const (
	DEFAULT_GRPC_PORT    = 8080
	DEFAULT_GRPC_TIMEOUT = "5s"
	DEFAULT_DEBUG        = false
)

type Config struct {
	Debug       bool   `koanf:"DASHBOARDS_DEBUG"`
	GrpcPort    int    `koanf:"PORT"`
	GrpcTimeout string `koanf:"DASHBOARDS_GRPC_TIMEOUT"`
}

func NewConfig() *Config {
	var c Config

	k := koanf.New(".")

	err := k.Load(confmap.Provider(map[string]interface{}{
		"DASHBOARDS_DEBUG":        DEFAULT_DEBUG,
		"PORT":                    DEFAULT_GRPC_PORT,
		"DASHBOARDS_GRPC_TIMEOUT": DEFAULT_GRPC_TIMEOUT,
	}, "."), nil)
	if err != nil {
		panic(err)
	}

	err = k.Load(env.Provider("DASHBOARDS_", ".", nil), nil)
	if err != nil {
		panic(err)
	}

	err = k.Unmarshal("", &c)
	if err != nil {
		panic(err)
	}

	return &c
}
