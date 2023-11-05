package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	key, err := os.ReadFile("key/ecdsa")
	if err != nil {
		log.Fatal(err)
	}

	manager, err := NewManager(key)
	if err != nil {
		log.Fatal(err)
	}

	srv := NewServer(manager)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	mux := http.NewServeMux()
	mux.Handle("/token", TokenHandler{srv})

	log.Fatal(http.ListenAndServe("localhost:9096", c.Handler(mux)))
}
