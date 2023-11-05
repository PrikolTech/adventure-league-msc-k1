package http

import (
	"auth-service/internal/service/oauth"
	"net/http"

	"github.com/rs/cors"
)

func NewServer(service *oauth.OAuth) *http.Server {
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux.Handle("/token", TokenHandler{service.Server})

	return &http.Server{
		Addr:    "localhost:9096",
		Handler: c.Handler(mux),
	}
}
