package internal

import (
	"context"
	"fmt"
	"net/http"
)

type (
	ctxKey int
)

const (
	ReverseProxyPublicUrl ctxKey = iota + 1001
)

// ReverseProxy this middleware adds the public url (i.e. reverse proxy) in the context
func ReverseProxy() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			proto := "http"
			if r.TLS != nil {
				proto = "https"
			}
			if value := r.Header.Get("X-Forwarded-Proto"); value != "" {
				proto = value
			}

			host := r.Host
			if value := r.Header.Get("X-Forwarded-Host"); value != "" {
				host = value
			}

			url := fmt.Sprintf("%s://%s", proto, host)
			ctx = context.WithValue(ctx, ReverseProxyPublicUrl, url)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
