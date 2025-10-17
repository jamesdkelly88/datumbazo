package main

import (
	"fmt"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/config"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/handlers"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/middleware"
	"net/http"
)

func main() {

	// define routes
	public := RouteList{
		"GET /favicon.ico":  handlers.FaviconHandler,
		"GET /health-check": handlers.HealthCheckHandler,
	}
	private := RouteList{
		"GET /version": handlers.VersionHandler,
	}
	v1routes := RouteList{
		"GET /": handlers.RootHandler1,
	}
	v2routes := RouteList{
		"GET /": handlers.RootHandler2,
	}

	// define main router
	router := http.NewServeMux()

	// define versioned api routers and prefixes
	v1 := http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1", v1))
	v2 := http.NewServeMux()
	router.Handle("/v2/", http.StripPrefix("/v2", v2))

	// define middleware stack (without auth)
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	// load routes
	addRoutes(router, stack, public, private)
	addRoutes(v1, stack, RouteList{}, v1routes)
	addRoutes(v2, stack, RouteList{}, v2routes)

	// define listen address
	address := fmt.Sprintf(":%d", config.Settings.Server.Port)

	// define server
	server := http.Server{
		Addr:    address,
		Handler: router,
	}
	// start server
	fmt.Printf("Starting server on %s\n", address)
	server.ListenAndServe()
}
