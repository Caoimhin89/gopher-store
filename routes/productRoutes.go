package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setProductRoutes(router *mux.Router) *mux.Router{
	for _, route := range productRoutes {
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

var productRoutes = Routes{

	Route{
		"GetProducts",
		"GET",
		"/products",
		GetProducts,
	},
	Route{
		"GetProduct",
		"GET",
		"/products/{id}",
		GetProduct,
	},
	Route{
		"AddProduct",
		"POST",
		"/products",
		AddProduct,
	},
	Route{
		"EditProduct",
		"PUT",
		"/products/{id}",
		EditProduct,
	},
	Route{
		"DeleteProduct",
		"DELETE",
		"/products/{id}",
		DeleteProduct,
	},

}