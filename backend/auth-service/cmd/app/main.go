package main

import (
	"auth-service/internal/app"
	"log"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run(cfg))
}
