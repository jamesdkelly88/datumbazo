package middleware

import (
	"errors"
	"log/slog"
	"net/http"
)

type User struct {
	Name   string
	Access string
}

func Authenticate(username string, password string) (string, error) {
	// TODO: check something for real users
	if username == "" {
		return "", errors.New("no username")
	}
	return "example", nil
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("BasicAuth middleware") // TODO: convert to trace log

		// get basic auth credentials
		var username, password, ok1 = r.BasicAuth()
		// get user from context
		user, ok2 := r.Context().Value(userKey).(*User)
		// authenticate user
		if ok1 && ok2 {
			access, err := Authenticate(username, password)
			if err == nil {
				user.Name = username
				user.Access = access
			}
		}
		// reject with 401 if user has no access
		if user.Access == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// pass to next handler
		next.ServeHTTP(w, r)
	})
}
