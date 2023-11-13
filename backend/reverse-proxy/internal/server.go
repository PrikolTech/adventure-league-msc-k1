package internal

import (
	"net/http"
)

type ServerOptions struct {
	Addr    string
	Origins []string
}

func NewServer(proxy ReverseProxy, opts ServerOptions) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", proxy.Handler())

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
