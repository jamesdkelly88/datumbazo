package handlers

import (
	"fmt"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/config"
	"net/http"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, config.Embedded, "favicon.ico")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"alive": true}`)
}

// func unmappedHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Unmapped path: %s", r.URL.RequestURI())
// }

func RootHandler1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Using v1 api")
}

func RootHandler2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Using v2 api")
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, config.Version.Full)
}
