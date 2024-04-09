package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func parseEnv() (*config, error) {
	conf := new(config)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, errors.Wrap(err, "failed to parse environment variables")
	}
	return conf, nil
}

type dbConfig struct {
	Host               string `envconfig:"host"`
	Name               string `envconfig:"name"`
	User               string `envconfig:"user"`
	Password           string `envconfig:"password"`
	MaxConn            int    `envconfig:"max_conn" default:"0"`
	ConnectionLifetime int    `envconfig:"conn_lifetime" default:"0"`
}

type serviceConfig struct {
	ServeRESTAddress string `envconfig:"serve_rest_address" default:"8082"`
}

type config struct {
	DB      dbConfig      `envconfig:"db"`
	Service serviceConfig `envconfig:"service"`
}
