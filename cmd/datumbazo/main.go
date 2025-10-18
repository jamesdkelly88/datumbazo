package main

import (
	"fmt"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/handlers"
	"github.com/jamesdkelly88/datumbazo/internal/config"
	"github.com/jamesdkelly88/datumbazo/internal/middleware"
	"net/http"
)

func main() {
	cfg := config.NewSettings(true)

	// define main router
	router := http.NewServeMux()

	// Create a base middleware chain
	baseChain := middleware.Chain{middleware.Logging}

	// Extend the base chain with basic auth
	authChain := append(baseChain, middleware.BasicAuth)

	// define routes
	router.Handle("GET /favicon.ico", baseChain.ThenFunc(handlers.FaviconHandler))
	router.Handle("GET /health-check", baseChain.ThenFunc(handlers.HealthCheckHandler))
	router.Handle("GET /version", authChain.ThenFunc(handlers.VersionHandler(cfg.Version)))
	router.Handle("/v1/", authChain.ThenFunc(handlers.RootHandler1))
	router.Handle("/v2/", authChain.ThenFunc(handlers.RootHandler2))

	// define listen address
	address := fmt.Sprintf(":%d", cfg.Server.Port)

	// define server
	server := http.Server{
		Addr:    address,
		Handler: router,
	}
	// start server
	fmt.Printf("Starting server on %s\n", address)
	server.ListenAndServe()
}
