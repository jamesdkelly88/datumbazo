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
	http.HandleFunc("/health-check", HealthCheckHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/", unmappedHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", settings.Server.Port), nil))
}
