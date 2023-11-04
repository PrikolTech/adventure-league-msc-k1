package http

import (
	"user-service/internal/service"
	"user-service/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type Services struct {
	User service.User
	Role service.Role
}

func Router(services Services) *chi.Mux {
	mux := chi.NewMux()

	user := handler.NewUser(services.User)

	mux.Route("/", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/", user.Get)
			r.Post("/", user.Create)
			r.Patch("/", user.Update)
			r.Delete("/", user.Delete)
			r.Post("/authenticate", user.Authenticate)
		})
		r.Route("/role", func(r chi.Router) {
		})
	})

	return mux
}
