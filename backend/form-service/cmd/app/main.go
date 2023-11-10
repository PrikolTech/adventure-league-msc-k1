package main

import (
	"form-service/internal/app"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()

	cfg, err := app.NewConfig()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	logger.Fatal().Msgf("%s", app.Run(cfg, &logger))
}
