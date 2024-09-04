package main

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 2)

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")
		if apiKey == "" {
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
