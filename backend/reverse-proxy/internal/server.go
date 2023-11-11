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
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   opts.Origins,
	// 	AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: true,
	// })

	mux.Handle("/", proxy.Handler())

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
