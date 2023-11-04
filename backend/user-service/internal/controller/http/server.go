package http

import (
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

func NewServer(logger zerolog.Logger, services Services, opts ServerOptions) *http.Server {
	mux := chi.NewMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   opts.Origins,
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux.Use(middleware.RealIP)
	mux.Use(LoggerMiddleware(logger))
	mux.Use(RecovererMiddleware(logger))
	mux.Use(c.Handler)
	mux.Use(ContentTypeMiddleware("application/json"))

	mux.NotFound(NotFound)
	mux.MethodNotAllowed(MethodNotAllowed)

	mux.Mount("/", Router(services))

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
