package middleware

import (
	"fmt"
	"github.com/jamesdkelly88/datumbazo/pkg/dbzo"
	"net/http"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, _, ok := r.BasicAuth()
		// username, password, ok := r.BasicAuth()
		if ok {
			fmt.Printf("Checking permissions for %s\n", username)
			user, err := dbzo.GetUser(username)
			if err == nil {
				fmt.Printf("User authorised - permissions are level %d\n", user.Permissions)
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
