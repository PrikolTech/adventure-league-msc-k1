package app

import (
	"auth-service/internal/net"
	"auth-service/internal/pkg/jwt"
	"auth-service/internal/service/oauth"
	"auth-service/internal/transport/http"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func Run(cfg *Config, logger *zerolog.Logger) error {
	privateKey, err := os.ReadFile(cfg.Key.Path)
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}

	user := net.NewUser(cfg.UserAPI.URL)
	role := net.NewRole(cfg.UserAPI.URL)

	parser, err := jwt.NewParserFromFile(cfg.Key.Path + ".pub")
	if err != nil {
		return fmt.Errorf("failed to parse public key: %w", err)
	}

	oauth, err := oauth.New(user, oauth.ManagerOptions{
		ClientID:     cfg.Client.ID,
		ClientDomain: cfg.Client.Domain,
		Key:          privateKey,
	})

	if err != nil {
		return fmt.Errorf("failed to init oauth: %w", err)
	}

	server := http.NewServer(http.Services{oauth, role, parser}, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info().Msg("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
