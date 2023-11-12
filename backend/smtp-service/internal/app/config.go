package app

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		HTTP ConfigHTTP
		SMTP ConfigSMTP
	}

	ConfigHTTP struct {
		Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
		Port string `env:"HTTP_PORT" env-default:"3009"`
	}

	ConfigSMTP struct {
		Host     string `env:"SMTP_HOST" env-required:""`
		Port     string `env:"SMTP_PORT" env-required:""`
		Username string `env:"SMTP_USERNAME" env-required:""`
		Password string `env:"SMTP_PASSWORD" env-required:""`
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
