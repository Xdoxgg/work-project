package main

import (
	"fmt"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Service: Hello, Product!")
}

func main() {
	fmt.Println("start")
	
	http.HandleFunc("/products", productHandler)
	fmt.Println("m1")
	http.ListenAndServe(":8082", nil)
	fmt.Println("m2")
	
}
