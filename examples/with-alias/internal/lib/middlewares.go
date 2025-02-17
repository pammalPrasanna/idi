package lib

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"strconv"
	"strings"

	"github.com/tomasen/realip"
)


func (i *rootApp) RequireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) == 2 && headerParts[0] == "Bearer" {
				token := headerParts[1]
				i.logger.Info(token)

				subject, err := i.VerifyToken(token)
				if err != nil {
					i.InvalidAuthenticationToken(w, r)
					return
				}
				userID, err := strconv.Atoi(subject)
				if err != nil {
					i.ServerError(w, r, err)
					return
				}

				user, err := i.GetUserById(userID)
				if err != nil {
					// decide based on the error
					i.ServerError(w, r, err)
					return
				}

				if user != nil {
					r = i.contextSetAuthenticatedUser(r, user)
				} else {
					i.AuthenticationRequired(w, r)
					return
				}
			}
		}

		// authenticatedUser := i.ContextGetAuthenticatedUser(r)

		// if authenticatedUser == nil {
		// 	i.AuthenticationRequired(w, r)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}


func (i *rootApp) recoverPanicM(next http.Handler) http.Handler {
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
func (i *rootApp) corsM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		if yes := slices.Contains(i.trustedOrigins, "*"); yes {
			// Use "*" for all origins, or replace with specific origins
			w.Header().Set("Access-Control-Allow-Origin", "*") 
			w.Header().Set("Access-Control-Allow-Credentials", "false")
		} else {
			rOrigin := r.Header.Get("Origin")
			if rOrigin != "" {
				if yes := slices.Contains(i.trustedOrigins, rOrigin); yes {
					w.Header().Set("Access-Control-Allow-Origin", rOrigin)
				}
			}
		}
		w.Header().Set("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Use "*" for all origins, or replace with specific origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (i *rootApp) loggerM(next http.Handler) http.Handler {
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
