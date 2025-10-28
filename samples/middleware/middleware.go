package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	r *http.Request
}

func logRequest(r *http.Request, start time.Time) {
	log.Printf("Request: method=%s, url=%s, user=%s, duration=%v", r.Method, r.URL.Path, r.Context().Value("username"), time.Since(start))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &LoggingResponseWriter{ResponseWriter: w, r: r}
		log.Println("Logging middleware inbound")
		log.Printf("%s\n", r.Context().Value("username"))
		next.ServeHTTP(lrw, r)
		log.Println("Logging middleware outbound")
		log.Printf("%s\n", r.Context().Value("username"))
		logRequest(lrw.r, start)
	})
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth middleware inbound")
		log.Printf("%s\n", r.Context().Value("username"))
		var username, _, ok = r.BasicAuth()
		if !ok || username == "" {
			username = "anonymous"
		}
		ctx := context.WithValue(r.Context(), "username", username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		log.Println("Auth middleware outbound")
		log.Printf("%s\n", r.Context().Value("username"))
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello handler start")
	username := r.Context().Value("username")
	if username == nil {
		username = "anonymous (at handler)"
	}
	fmt.Fprintf(w, "Hello, %s!\n", username)
	log.Println("Hello handler end")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", LoggingMiddleware(BasicAuthMiddleware(http.HandlerFunc(HelloHandler))))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
