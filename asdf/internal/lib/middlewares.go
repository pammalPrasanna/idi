package lib

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"github.com/tomasen/realip"
)

func (i *idi) recoverPanicM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				i.ServerError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORS middleware
func (i *idi) corsM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		if yes := slices.Contains(i.trustedOrigins, "*"); yes {
			w.Header().Set("Access-Control-Allow-Origin", "*") // Use "*" for all origins, or replace with specific origins
		} else {
			rOrigin := r.Header.Get("Origin")
			if rOrigin != "" {
				if yes := slices.Contains(i.trustedOrigins, rOrigin); yes {
					w.Header().Set("Access-Control-Allow-Origin", rOrigin)
				}
			}
		}
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are needed

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (i *idi) loggerM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := newMetricsResponseWriter(w)
		next.ServeHTTP(mw, r)

		var (
			ip     = realip.FromRequest(r)
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := slog.Group("user", "ip", ip)
		requestAttrs := slog.Group("request", "method", method, "url", url, "proto", proto)
		responseAttrs := slog.Group("repsonse", "status", mw.StatusCode, "size", mw.BytesCount)

		i.logger.Info("access", userAttrs, requestAttrs, responseAttrs)
	})
}
