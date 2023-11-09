package http

import (
	"auth-service/internal/service/oauth"
	"net/http"

	"github.com/rs/cors"
)

type ServerOptions struct {
	Addr    string
	Origins []string
}

func NewServer(service *oauth.OAuth, opts ServerOptions) *http.Server {
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   opts.Origins,
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux.Handle("/token", TokenHandler{service.Server})

	return &http.Server{
		Addr:    opts.Addr,
		Handler: c.Handler(mux),
	}
}
