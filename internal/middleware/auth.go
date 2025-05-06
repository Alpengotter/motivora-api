package middleware

import (
	"motivora-backend/internal/handlers"
	"net/http"
)

func Auth(next http.Handler) http.HandlerFunc {
	return handlers.AuthMiddleware(next)
}
