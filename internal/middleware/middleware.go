package middleware

import "net/http"

func ApplyMiddlewares(handler http.Handler) http.HandlerFunc {
	return CORSMiddleware(LoggingMiddleware(handler))
}
