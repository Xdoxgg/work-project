package main

import (
	"fmt"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://user-service:8081/api/users")
	if err != nil {
		http.Error(w, "Failed to call user service", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Product Service with user data")

	defer resp.Body.Close()
}

func main() {
	fmt.Println("start")

	http.HandleFunc("/api/products", productHandler)
	http.ListenAndServe(":8082", nil)

}
