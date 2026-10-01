package main

import (
	"log"
	"net/http"
	"pulseroute/internal/auth"
	"pulseroute/internal/limiter"
	"pulseroute/internal/proxy"
)

func main() {
	l := limiter.NewTokenBucket(100, 10)
	p := proxy.NewReverseProxy("http://backend:8080")
	a := auth.NewValidator("secret-token")

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if !a.Validate(r.Header.Get("Authorization")) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !l.Allow(r.RemoteAddr) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		p.ServeHTTP(w, r)
	})

	log.Println("PulseRoute Gateway listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
