package app

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		HTTP    ConfigHTTP
		UserAPI ConfigUserAPI
		Key     ConfigKey
	}

	ConfigHTTP struct {
		Host    string   `env:"HTTP_HOST" env-default:"0.0.0.0" `
		Port    string   `env:"HTTP_PORT" env-default:"3001"`
		Origins []string `env:"HTTP_ORIGINS" env-required:""`
	}

	ConfigUserAPI struct {
		URL string `env:"USER_API" env-required:""`
	}

	ConfigKey struct {
		Path string `env:"KEY_PATH" env-required:""`
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
