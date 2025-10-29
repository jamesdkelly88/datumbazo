package middleware

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type ctxKey uint8

const userKey ctxKey = 0
const requestKey ctxKey = 1

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
		slog.Debug("Logging middleware") // TODO: convert to trace log
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		user := new(User)
		var requestID uint64 = 0
		r = r.WithContext(context.WithValue(r.Context(), userKey, user))
		r = r.WithContext(context.WithValue(r.Context(), requestKey, &requestID))
		next.ServeHTTP(wrapped, r)
		if user.Name == "" {
			user.Name = "-"
		}
		if user.Access == "" {
			user.Access = "None"
		}
		slog.Info(
			"request",
			slog.Uint64("id", requestID),
			slog.Int("statusCode", wrapped.statusCode),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("user", user.Name),
			slog.String("access", user.Access),
			slog.Duration("duration", time.Since(start)),
		)
	})
}

func SetupLogger(w io.Writer, l slog.Level) {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: l.Level(),
	}))
	slog.SetDefault(logger)
}
