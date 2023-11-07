package app

import (
	"context"
	"fmt"
	"form-service/internal/repo"
	"form-service/internal/service"
	"form-service/internal/transport/http"
	"form-service/pkg/postgres"

	"github.com/rs/zerolog"
)

func Run(cfg *Config, logger *zerolog.Logger) error {
	postgres, err := postgres.New(context.Background(), cfg.Postgres.URI)
	if err != nil {
		return err
	}
	defer postgres.Close()
	logger.Info().Msg("database connection established")

	registrationRepo := repo.NewRegistrationPG(postgres)
	registrationService := service.NewRegistration(registrationRepo)

	server := http.NewServer(logger, registrationService, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info().Msg("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
