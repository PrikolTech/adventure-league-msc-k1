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
	role := handler.NewRole(services.Role)

	mux.Route("/", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}", user.Get)
			r.Post("/", user.Create)
			r.Patch("/{id}", user.Update)
			r.Delete("/{id}", user.Delete)
			r.Post("/authenticate", user.Authenticate)
		})
		r.Route("/role", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Post("/{user_id}", role.Append)
				r.Delete("/{user_id}", role.Remove)
			})
			r.Get("/", role.Get)
			r.Post("/", role.Create)
			r.Delete("/{id}", role.Delete)
		})
	})

	return mux
}