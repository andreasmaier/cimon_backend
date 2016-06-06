package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.Path("/ws").Handler(http.HandlerFunc(serveWs))

	return router
}

var routes = Routes{
	Route{
		"JobUpdates",
		"POST",
		"/jobupdates",
		JobUpdates,
	},
	Route{
		"WSTest",
		"POST",
		"/wstest",
		WSTest,
	},
}


