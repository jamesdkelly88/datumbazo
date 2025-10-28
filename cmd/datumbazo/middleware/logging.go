package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ctxKey uint8

const userKey ctxKey = 0

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Logging middleware") // TODO: convert to trace log
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		user := new(User)
		r = r.WithContext(context.WithValue(r.Context(), userKey, user))
		next.ServeHTTP(wrapped, r)
		if user.Name == "" {
			user.Name = "-"
		}
		if user.Access == "" {
			user.Access = "None"
		}
		fmt.Println(wrapped.statusCode, r.Method, r.URL.Path, user.Name, user.Access, time.Since(start))
	})
}
