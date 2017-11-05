package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setCollectionRoutes(router *mux.Router) *mux.Router{
	for _, route := range collectionRoutes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var collectionRoutes = Routes{

	Route{
		"GetCollections",
		"GET",
		"/collections",
		GetCollections,
	},
	Route{
		"GetCollection",
		"GET",
		"/collections/{id}",
		GetCollections,
	},
	Route{
		"AddCollection",
		"POST",
		"/collections",
		AddCollection,
	},
	Route{
		"EditCollection",
		"PUT",
		"/collections/{id}",
		EditCollection,
	},
	Route{
		"DeleteCollection",
		"DELETE",
		"/collections/{id}",
		DeleteCollection,
	},

}