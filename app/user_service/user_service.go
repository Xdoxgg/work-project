package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return nil, err
	}
	// Проверяем соединение
	if err = db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return nil, err
	}
	return db, nil
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
	db, err := connectDB()
	rows, err := db.Query("SELECT id, email, password FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
}

func postUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	_, err := connectDB()
	if err != nil {
		fmt.Println("conneting to db faled")
		return
	}
	http.HandleFunc("/api/users", handler)
	http.ListenAndServe(":8082", nil)
}
