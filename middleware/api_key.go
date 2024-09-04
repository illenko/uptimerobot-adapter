package middleware

import "net/http"

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")
		if apiKey == "" {
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
