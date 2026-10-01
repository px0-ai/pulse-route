package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewReverseProxy(targetURL string) *ReverseProxy {
	target, _ := url.Parse(targetURL)
	return &ReverseProxy{
		target: target,
		proxy:  httputil.NewSingleHostReverseProxy(target),
	}
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("X-Forwarded-By", "PulseRoute-v1")
	p.proxy.ServeHTTP(w, r)
}
