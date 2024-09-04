package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 2)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
