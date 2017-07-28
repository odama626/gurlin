package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter() //.StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		Index,
	},
	Route{
		"Available",
		"GET",
		"/api/available/{src}",
		Available,
	},
	Route{
		"Register",
		"POST",
		"/api/register",
		Register,
	},
	Route{
		"Delete",
		"DELETE",
		"/api/delete/{redirectKey}",
		Delete,
	},
	Route{
		"Disable",
		"POST",
		"/api/disable",
		Disable,
	},
	Route{
		"Bounce",
		"GET",
		"/{src}",
		Bounce,
	},
}
