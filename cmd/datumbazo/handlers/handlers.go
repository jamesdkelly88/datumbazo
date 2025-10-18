package handlers

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/jamesdkelly88/datumbazo/internal/config"
)

//go:embed favicon.ico
var Embedded embed.FS

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, Embedded, "favicon.ico")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"alive": true}`)
}

func UnmappedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unmapped path: %s", r.URL.RequestURI())
}

func RootHandler1(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Using v1 api")
}

func RootHandler2(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Using v2 api")
}

// func VersionHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, config.Version.Full)
// }

func VersionHandler(ver config.Version) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ver.Full)
	}
}
