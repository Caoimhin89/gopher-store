package main

import "github.com/gorilla/mux"

func setProductRoutes(router *mux.Router) *mux.Router{
	router.HandleFunc("/products", showProducts)
	router.HandleFunc("/products", addProduct)
	router.HandleFunc("/products/{id}", showProduct)
	router.HandleFunc("/products/{id}", updateProduct)
	router.HandleFunc("/products/{id}", deleteProduct)
	return router
}