package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setUserRoutes(router *mux.Router) *mux.Router{
	for _, route := range userRoutes {
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

var userRoutes = Routes{

	Route{
		"GetUsers",
		"GET",
		"/users",
		GetUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{id}",
		GetUser,
	},
	Route{
		"AddUser",
		"POST",
		"/users",
		AddUser,
	},
	Route{
		"EditUser",
		"PUT",
		"/users/{id}",
		EditUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/users/{id}",
		DeleteUser,
	},

}