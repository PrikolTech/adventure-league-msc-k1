package http

import (
	"form-service/internal/net"
	"form-service/internal/service"
	"form-service/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type Services struct {
	Registration service.Registration
	Auth         net.Auth
}

func Router(services Services) *chi.Mux {
	mux := chi.NewMux()

	registration := handler.NewRegistration(services.Registration)
	auth := handler.AuthMiddleware(services.Auth)

	mux.Route("/", func(r chi.Router) {
		r.Route("/registration", func(r chi.Router) {
			r.With(auth).Get("/{id}", registration.Get)
			r.With(auth).Get("/", registration.List)
			r.Post("/", registration.Create)
			r.With(auth).Post("/append", registration.Append)
			r.With(auth).Patch("/{id}", registration.Update)
		})
	})

	return mux
}
