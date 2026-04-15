package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Path        string
	Description string
	Controller  HandlerFunc
	Middleware  []Middleware
}

func (e *HttpEngine) AddRoutes(routes []Route) {
	for _, route := range routes {
		var handler http.Handler = http.HandlerFunc(route.Controller)

		for i := len(route.Middleware) - 1; i >= 0; i-- {
			handler = route.Middleware[i](handler)
		}

		finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := e.SetContext(r.Context())
			handler.ServeHTTP(w, r.WithContext(ctx))
		})

		fmt.Printf("%s - %s - %s\n", route.Path, route.Method, route.Description)
		e.router.Method(route.Method, route.Path, finalHandler)
	}
	fmt.Println()
}
