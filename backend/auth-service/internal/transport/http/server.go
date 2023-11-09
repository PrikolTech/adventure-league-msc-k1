package http

import (
	"auth-service/internal/net"
	"auth-service/internal/pkg/jwt"
	"auth-service/internal/service/oauth"
	"net/http"

	"github.com/rs/cors"
)

type ServerOptions struct {
	Addr    string
	Origins []string
}

type Services struct {
	OAuth  *oauth.OAuth
	Role   net.Role
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

	mux.Handle("/token", TokenHandler{services.OAuth.Server})
	mux.Handle("/verify", VerifyHandler{services.Role, services.Parser})

	return &http.Server{
		Addr:    opts.Addr,
		Handler: c.Handler(mux),
	}
}
