package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/products", showProducts)
	router.HandleFunc("/collections", showCollections)
	router.HandleFunc("/blog", showBlog)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func showProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products, %q", html.EscapeString(r.URL.Path))
}

func showCollections(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Collections, %q", html.EscapeString(r.URL.Path))
}

func showBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blog, %q", html.EscapeString(r.URL.Path))
}

