package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map[string]http.HandlerFunc
}

func (router *Router) Register(urlPath string, handler http.HandlerFunc) {
	router.handlers[urlPath] = handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := router.handlers[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Resource not found")
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]http.HandlerFunc)}
}

/*
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/products":
		fmt.Fprintln(w, "All product requests will be processed")
	case "/customers":
		fmt.Fprintln(w, "All customer requests will be processed")
	case "/":
		fmt.Fprintln(w, "Welcome to Go web server")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Welcome to Go web server")
	}
}
*/

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All product requests will be processed")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All customer requests will be processed")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go web server")
}

func main() {
	router := NewRouter()
	router.Register("/products", productsHandler)
	router.Register("/customers", customersHandler)
	router.Register("/", indexHandler)
	http.ListenAndServe(":8080", router)
}
