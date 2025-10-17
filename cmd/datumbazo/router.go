package main

import (
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/middleware"
	"net/http"
)

type RouteList map[string]http.HandlerFunc

func addRoute(router *http.ServeMux, stack middleware.Middleware, path string, handler http.HandlerFunc, authenticated bool) {
	if authenticated {
		router.Handle(path, middleware.BasicAuth(stack(http.HandlerFunc(handler))))
	} else {
		router.Handle(path, stack(http.HandlerFunc(handler)))
	}
}

func addRoutes(router *http.ServeMux, stack middleware.Middleware, public map[string]http.HandlerFunc, private map[string]http.HandlerFunc) {
	for publicPath, publicHandler := range public {
		router.Handle(publicPath, stack(http.HandlerFunc(publicHandler)))
	}
	for privatePath, privateHandler := range private {
		router.Handle(privatePath, middleware.BasicAuth(stack(http.HandlerFunc(privateHandler))))
	}
}
