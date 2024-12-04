package middleware

import (
	"net/http"
	"sample/internal/auth"
	"strings"
)

// JWTMiddleware is a middleware that checks if the request has a valid JWT token
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse and validate the token
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Attach the claims to the request context
		// You can also add authorization checks here based on the role
		r.Header.Set("username", claims.Username)
		r.Header.Set("role", claims.Role)

		next(w, r)
	}
}
