package internal

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type ReverseProxy interface {
	AddTarget(include []string, exclude []string, upstream string, prefix string) error
	Handler() http.Handler
}

type reverseProxy struct {
	proxy   *httputil.ReverseProxy
	targets []*Target
}

type Target struct {
	includes []*mux.Router
	excludes []*mux.Router
	upstream *url.URL
	prefix string
}

func NewReverseProxy() *reverseProxy {
	return &reverseProxy{
		proxy:   &httputil.ReverseProxy{},
		targets: make([]*Target, 0),
	}
}

func (p *reverseProxy) Rewrite() func(*httputil.ProxyRequest) {
	return func(pr *httputil.ProxyRequest) {
		for _, t := range p.targets {
			match := &mux.RouteMatch{}
			for _, exclude := range t.excludes {
				if exclude.Match(pr.In, match) {
					return
				}
			}
			for _, include := range t.includes {
				if include.Match(pr.In, match) {
					pr.Out.URL.Path = strings.TrimPrefix(pr.In.URL.Path, t.prefix)
					pr.SetURL(t.upstream)
					return
				}
			}
		}
	}
}

func (p *reverseProxy) AddTarget(include []string, exclude []string, upstream string, prefix string) error {
	url, err := url.Parse(upstream)
	if err != nil {
		return err
	}

	includeRouters := make([]*mux.Router, 0, len(include))
	excludeRouters := make([]*mux.Router, 0, len(exclude))

	for _, pattern := range include {
		router := mux.NewRouter()
		router.PathPrefix(prefix + pattern)
		includeRouters = append(includeRouters, router)
	}

	for _, pattern := range exclude {
		router := mux.NewRouter()
		router.PathPrefix(prefix + pattern)
		excludeRouters = append(excludeRouters, router)
	}

	p.targets = append(p.targets, &Target{
		includes: includeRouters,
		excludes: excludeRouters,
		upstream: url,
		prefix: prefix,
	})

	return nil
}

func (p *reverseProxy) Handler() http.Handler {
	p.proxy.Rewrite = p.Rewrite()
	return p.proxy
}
