package handlers

import (
	"embed"
	"fmt"
	"net/http"
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

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Version from config here") // TODO: get config in here
}
