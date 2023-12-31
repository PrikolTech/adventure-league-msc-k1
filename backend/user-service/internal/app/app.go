package app

import (
	"context"
	"fmt"
	"user-service/internal/net"
	"user-service/internal/repo"
	"user-service/internal/service"
	"user-service/internal/transport/http"
	"user-service/pkg/postgres"

	"github.com/rs/zerolog"
)

func Run(cfg *Config, logger *zerolog.Logger) error {
	postgres, err := postgres.New(context.Background(), cfg.Postgres.URI)
	if err != nil {
		return err
	}
	defer postgres.Close()
	logger.Info().Msg("database connection established")

	userRepo := repo.NewUserPG(postgres)
	roleRepo := repo.NewRolePG(postgres)

	userService := service.NewUser(userRepo, roleRepo)
	roleService := service.NewRole(roleRepo)

	authNet := net.NewAuth(cfg.AuthAPI.URL)

	server := http.NewServer(logger, http.Services{userService, roleService, authNet}, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info().Msg("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
