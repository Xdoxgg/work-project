// user_service.go
package main

import (
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsersHandler(w, r)
	case http.MethodPost:
		postUsersHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	
}

func postUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/api/users", handler)
	http.ListenAndServe(":8082", nil)
}
