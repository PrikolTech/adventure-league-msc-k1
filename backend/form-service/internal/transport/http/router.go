package http

import (
	"form-service/internal/service"

	"github.com/go-chi/chi/v5"
)

func Router(service service.Registration) *chi.Mux {
	mux := chi.NewMux()

	// registration := handler.NewRegistration(service)

	mux.Route("/", func(r chi.Router) {
		r.Route("/registration", func(r chi.Router) {
		})
	})

	return mux
}
