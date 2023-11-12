package http

import (
	"form-service/internal/transport/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

type ServerOptions struct {
	Addr    string
	Origins []string
}

func NewServer(logger *zerolog.Logger, services Services, opts ServerOptions) *http.Server {
	mux := chi.NewMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   opts.Origins,
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPatch},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.Use(middleware.RealIP)
	mux.Use(handler.LoggerMiddleware(logger))
	mux.Use(handler.RecovererMiddleware(logger))
	mux.Use(c.Handler)
	mux.Use(handler.ContentTypeMiddleware("application/json"))

	mux.NotFound(handler.NotFound)
	mux.MethodNotAllowed(handler.MethodNotAllowed)

	mux.Mount("/", Router(services))

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
