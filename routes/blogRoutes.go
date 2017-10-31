package main

import "github.com/gorilla/mux"

func setBlogRoutes(router *mux.Router) *mux.Router{
	for _, route := range blogRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var blogRoutes = Routes{

	Route{
		"GetBlogposts",
		"GET",
		"/blog",
		GetBlogposts,
	},
	Route{
		"GetBlogpost",
		"GET",
		"/blog/{id}",
		GetBlogpost,
	},
	Route{
		"AddBlogpost",
		"POST",
		"/blog",
		AddBlogpost,
	},
	Route{
		"EditBlogpost",
		"PUT",
		"/blog/{id}",
		EditBlogpost,
	},
	Route{
		"DeleteBlogpost",
		"DELETE",
		"/blog/{id}",
		DeleteBlogpost,
	},

}