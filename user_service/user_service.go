// user_service.go
package main

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Service: Hello, User!")
}

func main() {
	http.HandleFunc("/users", userHandler)
	http.ListenAndServe(":8081", nil)
}
