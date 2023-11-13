package main

import (
	"log"
	"smtp-service/internal/app"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run(cfg))
}
