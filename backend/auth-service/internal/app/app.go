package app

import (
	"auth-service/internal/net"
	"auth-service/internal/service/oauth"
	"auth-service/internal/transport/http"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func Run(cfg *Config, logger *zerolog.Logger) error {
	key, err := os.ReadFile(cfg.Key.Path)
	if err != nil {
		return err
	}

	user := net.NewUser(cfg.UserAPI.URL)

	service, err := oauth.New(user, oauth.ManagerOptions{
		ClientID:     cfg.Client.ID,
		ClientDomain: cfg.Client.Domain,
		Key:          key,
	})

	if err != nil {
		return err
	}

	server := http.NewServer(service, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info().Msg("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
