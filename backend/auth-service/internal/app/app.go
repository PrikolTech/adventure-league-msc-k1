package app

import (
	"auth-service/internal/net"
	"auth-service/internal/service/oauth"
	"auth-service/internal/transport/http"
	"fmt"
	"os"
)

func Run(cfg *Config) error {
	key, err := os.ReadFile(cfg.Key.Path)
	if err != nil {
		return err
	}

	user := net.NewUser(cfg.UserAPI.URL)

	service, err := oauth.New(user, key)
	if err != nil {
		return err
	}

	server := http.NewServer(service, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	return server.ListenAndServe()
}
