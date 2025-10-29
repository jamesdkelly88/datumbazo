package router

import (
	"fmt"
	"net/http"

	h "github.com/jamesdkelly88/datumbazo/cmd/datumbazo/handlers"
	mw "github.com/jamesdkelly88/datumbazo/cmd/datumbazo/middleware"
	cfg "github.com/jamesdkelly88/datumbazo/internal/config"
)

type Route struct {
	pattern       string
	handler       http.HandlerFunc
	authenticated bool
}

func Serve(listen string, settings cfg.Settings) {
	routes := []Route{
		{"/", h.Root, true},
		{"GET /favicon.ico", h.Favicon, false},
		{"GET /healthz", h.Health, false},
		{"GET /version", h.Version(settings.Version), true},
	}

	router := http.NewServeMux()

	publicChain := mw.Chain{mw.Logging}
	authChain := mw.Chain{mw.Logging, mw.BasicAuth}

	for _, r := range routes {
		if r.authenticated {
			router.Handle(r.pattern, authChain.ThenFunc(r.handler))
		} else {
			router.Handle(r.pattern, publicChain.ThenFunc(r.handler))
		}
	}

	server := http.Server{
		Addr:    listen,
		Handler: router,
	}
	// start server
	fmt.Printf("Starting server on %s\n", server.Addr) // TODO: central logging
	server.ListenAndServe()
}
