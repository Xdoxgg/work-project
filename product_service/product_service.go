package main

import (
	"fmt"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Service: Hello, Product!")
}

func main() {
	http.HandleFunc("/products", productHandler)
	http.ListenAndServe(":8082", nil)
}
