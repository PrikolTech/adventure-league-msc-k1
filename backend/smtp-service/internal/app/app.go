package app

import (
	"fmt"
	"log"
	"smtp-service/internal/service"
	"smtp-service/internal/transport/http"
)

func Run(cfg *Config) error {
	sender := service.NewSender(&service.SenderOptions{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		Username: cfg.SMTP.Username,
		Password: cfg.SMTP.Password,
	})

	server := http.NewServer(sender, &http.ServerOptions{
		Addr: fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
	})

	log.Printf("server created with address %s\n", server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
