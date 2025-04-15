package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

func connectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL") // Используем переменную окружения

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return nil, err
	}
	return db, nil
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMoviesHandler(w, r)
	case http.MethodPost:
		postMovieHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description, year FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func postMovieHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")
	year := r.URL.Query().Get("year")

	if title == "" || description == "" || year == "" {
		http.Error(w, "Все поля обязательны", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO movies (title, description, year) VALUES ($1, $2, $3)",
		title, description, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	port := os.Getenv("MOVIE_SERVICE_PORT") // Используем переменную окружения для порта
	if port == "" {
		port = "8080" // Значение по умолчанию
	}

	fmt.Println("movie-service started on port", port)
	http.HandleFunc("/api/movies", movieHandler)
	http.ListenAndServe(":"+port, nil)
}
