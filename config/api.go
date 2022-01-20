package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Api struct {
	Name            string        `default:"tracker"`
	Host            string        `default:"0.0.0.0"`
	Port            string        `default:"5000"`
	Mode            string        `default:"dev"`
	LogLevel        string        `default:"debug"`
	GracefulTimeout time.Duration `default:"8s"`
}

func API() Api {
	var api Api
	envconfig.MustProcess("API", &api)

	return api
}
