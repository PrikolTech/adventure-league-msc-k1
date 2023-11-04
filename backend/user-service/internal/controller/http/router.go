package http

import (
	"user-service/internal/service"

	"github.com/go-chi/chi/v5"
)

type Services struct {
	User service.User
	Role service.Role
}

func Router(services Services) *chi.Mux {
	mux := chi.NewMux()

	mux.Route("/", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
		})
		r.Route("/role", func(r chi.Router) {
		})
	})

	return mux
}
