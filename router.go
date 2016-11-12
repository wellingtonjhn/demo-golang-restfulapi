package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHttpRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		LogRoute(&route)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func LogRoute(route *Route) {
	var handler http.Handler
	handler = route.HandlerFunc
	handler = Logger(handler, route.Name)
}
