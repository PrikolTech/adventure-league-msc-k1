package http

import (
	"net/http"
	"smtp-service/internal/service"
	"smtp-service/internal/transport/http/handler"
)

type ServerOptions struct {
	Addr string
}

func NewServer(sender service.Sender, opts *ServerOptions) *http.Server {
	mux := http.NewServeMux()

	password := handler.NewPassword(sender)
	mux.Handle("/password", password)

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
