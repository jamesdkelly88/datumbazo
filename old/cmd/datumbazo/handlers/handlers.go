package handlers

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jamesdkelly88/datumbazo/internal/config"
)

//go:embed favicon.ico
var Embedded embed.FS

func Favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, Embedded, "favicon.ico")
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Welcome to Datumbazo")
	}
}

func Version(ver config.Version) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(ver)
		if err == nil {
			fmt.Fprint(w, string(bytes))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
