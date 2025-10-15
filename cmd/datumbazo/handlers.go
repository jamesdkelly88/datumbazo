package main

import (
	"fmt"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, embedded, "favicon.ico")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"alive": true}`)
}

func unmappedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unmapped path: %s", r.URL.RequestURI())
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, version.Full)
}
