package middleware

import (
	"context"
	// "fmt"
	"github.com/jamesdkelly88/datumbazo/internal/user"
	"net/http"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, _, ok := r.BasicAuth()
		// username, password, ok := r.BasicAuth()
		if ok {
			// fmt.Printf("Checking permissions for %s\n", username)
			_, err := user.GetUser(username)
			// user, err := user.GetUser(username)
			if err == nil {
				// fmt.Printf("User authorised - permissions are level %d\n", user.Permissions)
				ctx := context.WithValue(r.Context(), "UserId", username)
				req := r.WithContext(ctx)
				next.ServeHTTP(w, req)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
