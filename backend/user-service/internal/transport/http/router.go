package http

import (
	"user-service/internal/net"
	"user-service/internal/service"
	"user-service/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type Services struct {
	User service.User
	Role service.Role
	Auth net.Auth
}

func Router(services Services) *chi.Mux {
	mux := chi.NewMux()

	user := handler.NewUser(services.User)
	role := handler.NewRole(services.Role)

	auth := handler.AuthMiddleware(services.Auth)

	mux.Route("/", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/exist", user.Exist)
			r.Post("/authenticate", user.Authenticate)

			r.With(auth).Get("/{id}", user.Get)
			r.With(auth).Get("/", user.List)
			r.Post("/", user.Create)
			r.With(auth).Patch("/{id}", user.Update)
			r.With(auth).Delete("/{id}", user.Delete)
		})
		r.With(auth).Route("/role", func(r chi.Router) {
			r.Get("/", role.Get)
			r.Post("/", role.Create)
			r.Post("/append", role.Append)
			r.Delete("/{id}", role.Delete)
			r.Delete("/remove", role.Remove)
		})
	})

	return mux
}
