package config

import (
	"errors"
	"os"
)

type Config struct {
	Name string
	Port string
}

func LoadFromEnv() (*Config, error) {
	cfg := new(Config)

	env := os.Getenv("APP_NAME")
	if env == "" {
		return nil, errors.New("empty env: APP_NAME")
	}
	cfg.Name = env

	env = os.Getenv("HTTP_PORT")
	if env == "" {
		return nil, errors.New("empty env: HTTP_PORT")
	}
	cfg.Port = env

	return cfg, nil
}
