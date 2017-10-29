package main

import "github.com/gorilla/mux"

func setUserRoutes(router *mux.Router) *mux.Router{
	router.HandleFunc("/users", showUsers)
	router.HandleFunc("/users", addUser)
	router.HandleFunc("/users/{id}", showUser)
	router.HandleFunc("/users/{id}", updateUser)
	router.HandleFunc("/users/{id}", deleteUser)
	return router
}