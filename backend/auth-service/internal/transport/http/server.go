package http

import (
	"auth-service/internal/pkg/jwt"
	"auth-service/internal/service"
	"net/http"

	"github.com/rs/cors"
)

type ServerOptions struct {
	Addr    string
	Origins []string
}

type Services struct {
	OAuth  service.OAuth
	Parser jwt.Parser
}

func NewServer(services Services, opts ServerOptions) *http.Server {
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   opts.Origins,
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux.Handle("/token", TokenHandler{services.OAuth})
	mux.Handle("/verify", VerifyHandler{services.Parser})

	return &http.Server{
		Addr:    opts.Addr,
		Handler: c.Handler(mux),
	}
}
