package main

import (
	"log"
	"net/http"
)

func (s *server) APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")

		log.Println("Used API Key")

		if key == s.apiKey {
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func (s *server) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		log.Println("Used BasicAuth")

		if !ok || user != s.user || pass != s.pass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Log in to Thing API"`)
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (s *server) CustomAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Many possibilities

		// 1. Whenever a user logs in, add a session token to the cookies and
		// check the validity of the token here

		// 2. Use an external service to provide authentication, for example Auth0
		// https://auth0.com/blog/authentication-in-golang/

		// 3. Use self-issued JWT tokens
		// https://www.sohamkamani.com/golang/jwt-authentication/
	})
}
