package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[REQUEST] %s %s | RemoteAddr: %s | User-Agent: %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			r.UserAgent(),
		)
		next.ServeHTTP(w, r)
	}
}
