package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
	}

	HTTP struct {
		Port  string `env-required:"true" yaml:"port" env:"DAUNRODO_HTTP_PORT"`
		Proxy string `yaml:"proxy" env:"DAUNRODO_HTTP_PROXY"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"DAUNRODO_LOG_LEVEL"`
	}
)

// New returns app config.
func New(configPath string) (*Config, error) {

	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
