package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/jamesdkelly88/datumbazo/pkg/dbzo"
)

//go:embed favicon.ico

var embedded embed.FS
var settings dbzo.Settings
var version dbzo.Version

func main() {
	settings = dbzo.NewSettings()
	version = dbzo.GetVersion(true)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/", unmappedHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", settings.Server.Port), nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, embedded, "favicon.ico")
}

func unmappedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unmapped path: %s", r.URL.RequestURI())
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, version.Full)
}
