package app

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		Postgres ConfigPostgres
		HTTP     ConfigHTTP
	}

	ConfigPostgres struct {
		URI string `env:"POSTGRES_URI" env-required:""`
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
