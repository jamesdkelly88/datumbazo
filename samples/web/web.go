package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func query(w http.ResponseWriter, r *http.Request) {
	// check r.Method
	// read the body if present r.Body
	// check for query arguments r.URL.Query()
}

func unmapped(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unmapped path: %s", r.URL.RequestURI())
}

func version(w http.ResponseWriter, r *http.Request) {
	var name string = "TestApp"
	var version string = "0.1"

	fmt.Fprintf(w, "%s %s on %s %s, running %s, compiled by %s", name, version, runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.Compiler)
}

func main() {
	http.HandleFunc("/version", version)
	http.HandleFunc("/query", query)
	http.HandleFunc("/", unmapped)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
