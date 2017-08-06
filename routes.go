package main

import (
	"fmt"
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
	router := mux.NewRouter()    //.StrictSlash(true)
	apiRouter := mux.NewRouter() //	router.PathPrefix("/" + Env.API_ROUTE).Subrouter()

	router.PathPrefix("/" + Env.API_ROUTE).
		Handler(http.StripPrefix("/"+Env.API_ROUTE, adapt(apiRouter, authMiddleware)))

	for _, route := range routes {
		apiRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.HandleFunc("/{src}", Bounce)

	return router
}

func adapt(h http.Handler, routes ...func(http.Handler) http.Handler) http.Handler {
	for _, route := range routes {
		h = route(h)
	}
	return h
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Auth middleware")
		h.ServeHTTP(res, req)
	})
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"",
		Index,
	},
	Route{
		"Available",
		"GET",
		"/available/{src}",
		Available,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
	Route{
		"Delete",
		"DELETE",
		"/delete/{redirectKey}",
		Delete,
	},
	Route{
		"Disable",
		"POST",
		"/disable",
		Disable,
	},
}
