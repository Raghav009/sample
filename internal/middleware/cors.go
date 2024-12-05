package middleware

import (
	"net/http"
)

// CORS middleware to add CORS headers to the response
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4201/")           // Or set a specific domain instead of "*"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allowed headers
		w.Header().Set("Access-Control-Allow-Credentials", "true")                        // Allow credentials (cookies, etc.)

		// Handle preflight request (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
