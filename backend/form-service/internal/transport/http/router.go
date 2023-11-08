package http

import (
	"form-service/internal/service"
	"form-service/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

func Router(service service.Registration) *chi.Mux {
	mux := chi.NewMux()

	registration := handler.NewRegistration(service)

	mux.Route("/", func(r chi.Router) {
		r.Route("/registration", func(r chi.Router) {
			r.Get("/{id}", registration.Get)
			r.Get("/", registration.List)
			r.Post("/", registration.Create)
			r.Patch("/{id}", registration.Update)
		})
	})

	return mux
}
