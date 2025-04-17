package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Genre struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Recommendation struct {
	MovieID int `json:"movie_id"`
}

func connectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
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

func recomendationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRecommendationsHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Получаем жанры и теги из запроса
	var tags []Tag
	var genres []Genre

	// Измените здесь, чтобы использовать значения, а не указатели
	var requestBody struct {
		Tags   []Tag   `json:"tags"`
		Genres []Genre `json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}

	tags = requestBody.Tags
	genres = requestBody.Genres
	fmt.Println(tags, genres)
	// Получаем рекомендации
	recommendations, err := getRecommendations(db, tags, genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем рекомендации
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// Функция для получения рекомендаций на основе жанров и тегов
func getRecommendations(db *sql.DB, tags []Tag, genres []Genre) ([]Recommendation, error) {
	var movieIDs []int

	// Шаг 1: Найти все ID фильмов с указанными жанрами
	for _, genre := range genres {
		var ids []int
		query := "SELECT movie_id FROM movies_to_genres WHERE genre_id = $1"
		rows, err := db.Query(query, genre.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}
		movieIDs = append(movieIDs, ids...)
	}

	// Шаг 2: Найти все фильмы с такими же жанрами или тегами
	var recommendations []Recommendation
	for _, tag := range tags {
		query := "SELECT movie_id FROM movies_to_tags WHERE tag_id = $1"
		rows, err := db.Query(query, tag.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var movieID int
			if err := rows.Scan(&movieID); err != nil {
				return nil, err
			}
			// Проверяем, чтобы ID фильма уже не был добавлен
			if contains(movieIDs, movieID) {
				recommendations = append(recommendations, Recommendation{MovieID: movieID})
			}
		}
	}

	return recommendations, nil
}

// Функция для проверки наличия элемента в срезе
func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("recomendations_service started")
	http.HandleFunc("/api/recomendations", recomendationsHandler)
	http.ListenAndServe(":8080", nil)
}
