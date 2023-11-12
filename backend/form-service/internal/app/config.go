package app

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		Postgres  ConfigPostgres
		HTTP      ConfigHTTP
		UserAPI   ConfigUserAPI
		CourseAPI ConfigCourseAPI
		AuthAPI   ConfigAuthAPI
	}

	ConfigPostgres struct {
		URI string `env:"POSTGRES_URI" env-required:""`
	}

	ConfigHTTP struct {
		Host    string   `env:"HTTP_HOST" env-default:"0.0.0.0"`
		Port    string   `env:"HTTP_PORT" env-default:"3007"`
		Origins []string `env:"HTTP_ORIGINS" env-required:""`
	}

	ConfigUserAPI struct {
		URL string `env:"USER_API" env-required:""`
	}

	ConfigCourseAPI struct {
		URL string `env:"COURSE_API" env-required:""`
	}

	ConfigAuthAPI struct {
		URL string `env:"AUTH_API" env-required:""`
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
