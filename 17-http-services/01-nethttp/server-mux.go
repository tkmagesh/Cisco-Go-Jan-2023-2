package main

import (
	"fmt"
	"net/http"
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

func main() {

	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/customers", customersHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
