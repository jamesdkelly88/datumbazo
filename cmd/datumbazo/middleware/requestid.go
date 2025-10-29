package middleware

import (
	"log/slog"
	"net/http"
)

var nextId uint64 = 1

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("RequestID middleware") // TODO: convert to trace log
		id, ok := r.Context().Value(requestKey).(*uint64)
		if ok {
			*id = nextId
			// slog.Debug(fmt.Sprintf("RequestID set to %d", *id)) // TODO: convert to trace log
			nextId += 1
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// pass to next handler
		next.ServeHTTP(w, r)
	})
}
