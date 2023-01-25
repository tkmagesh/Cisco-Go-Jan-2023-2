package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All product requests will be processed")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All customer requests will be processed")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go web server")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Product -id[%v] will be processed", vars["id"])
}

/* middlewares */
type Middleware func(http.HandlerFunc) http.HandlerFunc

func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		log.Println("Time Taken : ", elapsed)
	}
}

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	router := mux.NewRouter()
	/*
		router.HandleFunc("/products", profileMiddleware(logMiddleware(productsHandler)))
		router.HandleFunc("/products/{id}", profileMiddleware(logMiddleware(productHandler)))
		router.HandleFunc("/customers", profileMiddleware(logMiddleware(customersHandler)))
		router.HandleFunc("/", profileMiddleware(logMiddleware(indexHandler)))
	*/

	router.HandleFunc("/products", chain(productsHandler, logMiddleware, profileMiddleware))
	router.HandleFunc("/products/{id}", chain(productHandler, logMiddleware, profileMiddleware))
	router.HandleFunc("/customers", chain(customersHandler, logMiddleware, profileMiddleware))
	router.HandleFunc("/", chain(indexHandler, logMiddleware, profileMiddleware))
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
