package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware 1: Authentication (check for valid token)
func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for Authorization header (e.g., Bearer token)
		username, _, ok := r.BasicAuth()
		if ok {
			// Extract the token (in a real app, validate the token)
			// Token is valid, add user info to context (if needed)
			r.Header.Set("X-User-Name", username) // Set a dummy user for the example
			next.ServeHTTP(w, r)                  // Proceed to the next middleware/handler
			return
		}
		// If no valid token, proceed with the request but let the logging and authorization happen
		next.ServeHTTP(w, r)
	})
}

// Middleware 2: Logging (log all requests)
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time for response time calculation
		start := time.Now()

		// Log the request method, URL, and user (from a custom header)
		username := r.Header.Get("X-User-Name") // Username should be set by authenticate middleware
		log.Printf("Request by user: %s, Method: %s, URL: %s", username, r.Method, r.URL.Path)

		// Proceed with the request
		next.ServeHTTP(w, r)

		// Calculate response time
		duration := time.Since(start)
		log.Printf("Response for %s took %s", r.URL.Path, duration)
	})
}

// Middleware 3: Authorization (check user permissions)
func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example authorization check (custom logic based on user)
		username := r.Header.Get("X-User-Name")
		if username == "admin" {
			// Allow admin to access all endpoints
			next.ServeHTTP(w, r)
		} else if r.URL.Path == "/admin" {
			// If not admin, block access to the admin route
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			// Allow access to non-admin endpoints
			next.ServeHTTP(w, r)
		}
	})
}

// Middleware 4: Error Handling (handle errors and panics)
func errorHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer function to catch panics
		defer func() {
			if err := recover(); err != nil {
				// Log the panic error details
				log.Printf("Error: %v", err)
				// Send 500 Internal Server Error response
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Proceed with the request
		next.ServeHTTP(w, r)
	})
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Example of normal behavior
	fmt.Fprintln(w, "Welcome to the home page!")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// Example of error behavior
	// This could simulate a panic or some internal error
	if r.URL.Query().Get("error") == "true" {
		panic("Something went wrong in the admin page!")
	}
	fmt.Fprintln(w, "Welcome to the admin page!")
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define routes and handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/admin", adminHandler)

	// Chain middlewares (authentication first, then logging, then authorization, then error handling)
	handler := authenticate(logRequest(authorize(errorHandling(mux))))

	// Start the server
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
