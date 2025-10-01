package middleware

import (
	"log"
	"net/http"
	"time"
)

// CorsMiddleware handles CORS for all requests including preflight
func CorsMiddleware(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}

// LoggingMiddleware wraps the mux with request logging
func LoggingMiddleware(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log the incoming request
		log.Printf("[%s] %s %s - Started", r.Method, r.URL.Path, r.RemoteAddr)
		
		// Call the mux
		mux.ServeHTTP(w, r)
		
		// Log the completed request with duration
		duration := time.Since(start)
		log.Printf("[%s] %s %s - Completed in %v", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}

// ChainMiddleware combines multiple middleware functions
func ChainMiddleware(mux *http.ServeMux, middlewares ...func(*http.ServeMux) http.Handler) http.Handler {
	handler := http.Handler(mux)
	
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](http.NewServeMux())
	}
	
	return handler
}

// RateLimiter (basic implementation - you can enhance this)
func RateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Basic rate limiting logic could go here
		// For now, just pass through
		next.ServeHTTP(w, r)
	}
}

// ValidateContentType ensures JSON content type for POST/PUT requests
func ValidateContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" {
			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}