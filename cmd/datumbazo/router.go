package main

import (
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/middleware"
	"net/http"
)

type RouteList map[string]http.HandlerFunc

func addRoute(router *http.ServeMux, stack middleware.Middleware, path string, handler http.HandlerFunc, authenticated bool) {
	if authenticated {
		router.Handle(path, stack(middleware.BasicAuth(http.HandlerFunc(handler))))
	} else {
		router.Handle(path, stack(http.HandlerFunc(handler)))
	}
}

func addRoutes(router *http.ServeMux, stack middleware.Middleware, public map[string]http.HandlerFunc, private map[string]http.HandlerFunc) {
	for publicPath, publicHandler := range public {
		router.Handle(publicPath, stack(http.HandlerFunc(publicHandler)))
	}
	for privatePath, privateHandler := range private {
		router.Handle(privatePath, stack(middleware.BasicAuth(http.HandlerFunc(privateHandler))))
	}
}

// func loadPublicRoutes(router ...*http.ServeMux) {
// 	stack := middleware.CreateStack(
// 		middleware.Logging,
// 	)
// 	for _, r := range router {
// 		r.Handle("GET /favicon.ico", stack(http.HandlerFunc(handlers.FaviconHandler)))
// 		r.Handle("GET /health-check", stack(http.HandlerFunc(handlers.HealthCheckHandler)))
// 	}
// }

// func loadPrivateRoutes(router ...*http.ServeMux) {

// 	// define middleware stack
// 	stack := middleware.CreateStack(
// 		middleware.Logging,
// 		middleware.BasicAuth,
// 	)
// 	for _, r := range router {
// 		r.Handle("GET /version", stack(http.HandlerFunc(handlers.VersionHandler)))
// 		r.Handle("/", stack(http.HandlerFunc(handlers.RootHandler)))
// 	}
// }
