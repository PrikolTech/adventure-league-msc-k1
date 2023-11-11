package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

type ReverseProxy struct {
	proxy   *httputil.ReverseProxy
	targets []*Target
}

type Target struct {
	includes []*mux.Router
	excludes []*mux.Router
	upstream *url.URL
}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{
		proxy:   &httputil.ReverseProxy{},
		targets: make([]*Target, 0),
	}
}

func (p *ReverseProxy) Rewrite() func(*httputil.ProxyRequest) {
	return func(pr *httputil.ProxyRequest) {
		fmt.Println("accepted", pr.In.Method, pr.In.URL.Path)
		for _, t := range p.targets {
			match := &mux.RouteMatch{}
			for _, exclude := range t.excludes {
				if exclude.Match(pr.In, match) {
					return
				}
			}
			for _, include := range t.includes {
				if include.Match(pr.In, match) {
					pr.SetURL(t.upstream)
					return
				}
			}
		}
	}
}

func (p *ReverseProxy) AddTarget(include []string, exclude []string, upstream string) error {
	url, err := url.Parse(upstream)
	if err != nil {
		return err
	}

	includeRouters := make([]*mux.Router, 0, len(include))
	excludeRouters := make([]*mux.Router, 0, len(exclude))

	for _, pattern := range include {
		router := mux.NewRouter()
		router.PathPrefix(pattern)
		includeRouters = append(includeRouters, router)
	}

	for _, pattern := range exclude {
		router := mux.NewRouter()
		router.PathPrefix(pattern)
		excludeRouters = append(excludeRouters, router)
	}

	p.targets = append(p.targets, &Target{
		includes: includeRouters,
		excludes: excludeRouters,
		upstream: url,
	})

	return nil
}

func (p *ReverseProxy) Start() {
	p.proxy.Rewrite = p.Rewrite()
	http.Handle("/", p.proxy)
}

func main() {
	proxy := NewReverseProxy()
	proxy.AddTarget([]string{"/token"}, nil, "http://localhost:3001")
	proxy.AddTarget([]string{"/user"}, []string{"/user/exist"}, "http://localhost:3002")
	proxy.AddTarget([]string{"/form"}, nil, "http://localhost:3007")

	proxy.Start()
	http.ListenAndServe("localhost:3000", nil)
}
