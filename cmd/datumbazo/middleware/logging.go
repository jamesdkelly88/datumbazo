package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// TODO: with the basicauth middleware on the outside of the stack, any 401s and / 404s are not getting logged
//       so need to remove the username context check from here and use it in the real handlers when checking permissions
//       this function should just be able to get the username from the header same as the BasicAuth middleware does

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		userId, ok := r.Context().Value("UserId").(string)

		if !ok {
			userId = "-"
		}

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		fmt.Println(wrapped.statusCode, r.Method, r.URL.Path, userId, time.Since(start))
	})
}
