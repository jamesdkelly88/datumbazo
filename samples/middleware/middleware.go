package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name string
}

func logRequest(r *http.Request, start time.Time, username string) {
	log.Printf("Request: method=%s, url=%s, user=%s, duration=%v", r.Method, r.URL.Path, username, time.Since(start))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Logging middleware inbound")
		user := new(User)
		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		next.ServeHTTP(w, r)
		log.Println("Logging middleware outbound")
		log.Printf("%s\n", user.Name)
		logRequest(r, start, user.Name)
	})
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth middleware inbound")
		log.Printf("%v\n", r.Context().Value("user"))
		var username, _, ok = r.BasicAuth()
		if !ok || username == "" {
			username = "anonymous"
		}
		user, ok := r.Context().Value("user").(*User)
		if ok {
			user.Name = username
		}
		next.ServeHTTP(w, r)
		log.Println("Auth middleware outbound")
		log.Printf("%v\n", r.Context().Value("user"))
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello handler start")
	var username string
	user, ok := r.Context().Value("user").(*User)
	if ok {
		username = user.Name
	} else {
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
