package main

import (
	"fmt"
	"log"
	"reverse-proxy/internal"
)

func main() {
	cfg, err := internal.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	targets, err := internal.ParseServices("configs/services.json")
	if err != nil {
		log.Fatal(err)
	}

	proxy := internal.NewReverseProxy()
	for _, target := range targets {
		proxy.AddTarget(target.Includes, target.Excludes, target.Upstream)
	}

	server := internal.NewServer(proxy, internal.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	log.Printf("server created with address %s\n", server.Addr)
	log.Fatalf("server down: %s", server.ListenAndServe())
}
