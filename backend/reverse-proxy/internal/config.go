package internal

import (
	"encoding/json"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP ConfigHTTP
	}

	ConfigHTTP struct {
		Host    string   `env:"HTTP_HOST" env-default:"0.0.0.0" `
		Port    string   `env:"HTTP_PORT" env-default:"3000"`
		Origins []string `env:"HTTP_ORIGINS" env-required:""`
	}
)

func NewConfig() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type TargetConfig struct {
	Upstream string   `json:"upstream"`
	Includes []string `json:"includes"`
	Excludes []string `json:"excludes"`
}

func ParseServices(path string) ([]TargetConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var targets []TargetConfig

	d := json.NewDecoder(file)
	err = d.Decode(&targets)
	return targets, err
}
