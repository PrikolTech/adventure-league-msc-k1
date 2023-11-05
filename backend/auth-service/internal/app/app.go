package app

import (
	"auth-service/internal/service/oauth"
	"auth-service/internal/transport/http"
	"auth-service/internal/webapi"
	"os"
)

func Run() error {
	key, err := os.ReadFile("key/ecdsa")
	if err != nil {
		return err
	}

	user := webapi.NewUser()

	service, err := oauth.New(user, key)
	if err != nil {
		return err
	}

	server := http.NewServer(service)

	return server.ListenAndServe()
}
