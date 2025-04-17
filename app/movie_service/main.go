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

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        string  `json:"year"`
	Genres      []Genre `json:"genres"`
	Tags        []Tag   `json:"tags"`
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
	//case http.MethodPost:
	//	postMovieHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func moviesGenresHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMoviesGenresHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getMoviesGenresHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	genres, err := getGenresTags(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}
func getGenresTags(db *sql.DB) ([]Genre, error) {
	rows, err := db.Query("SELECT * FROM genres ORDER BY title ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var genres []Genre
	for rows.Next() {
		var genre Genre
		if err := rows.Scan(&genre.ID, &genre.Title); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func moviesTagsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMoviesTagsHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getMoviesTagsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	tags, err := getMoviesTags(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func getMoviesTags(db *sql.DB) ([]Tag, error) {

	rows, err := db.Query("SELECT * FROM tags ORDER BY title ASC")

	if err != nil {
		return nil, err
	}
	defer db.Close()
	var tags []Tag
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Title); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	movies, err := getMovies(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovies(db *sql.DB) ([]Movie, error) {
	rows, err := db.Query("SELECT id, title, description, year FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year); err != nil {
			return nil, err
		}

		genresRows, err := db.Query("SELECT genres.id, genres.title FROM movies_to_genres JOIN genres ON (genre_id = genres.id) WHERE movie_id = $1", movie.ID)
		if err != nil {
			return nil, err
		}
		defer db.Close()

		var genres []Genre
		for genresRows.Next() {
			var genre Genre
			err := genresRows.Scan(&genre.ID, &genre.Title)
			if err != nil {
				return nil, err
			}
			genres = append(genres, genre)
		}
		movie.Genres = genres

		tagsRows, err := db.Query("SELECT tags.id, tags.title FROM movies_to_tags JOIN tags ON (tag_id = tags.id) WHERE movie_id = $1", movie.ID)
		if err != nil {
			return nil, err
		}
		defer db.Close()

		var tags []Tag
		for tagsRows.Next() {
			var tag Tag
			err := tagsRows.Scan(&tag.ID, &tag.Title)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		movie.Tags = tags

		movies = append(movies, movie)
	}
	return movies, nil
}

//func postMovieHandler(w http.ResponseWriter, r *http.Request) {
//	db, err := connectDB()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	defer db.Close()
//
//	title := r.URL.Query().Get("title")
//	description := r.URL.Query().Get("description")
//	year := r.URL.Query().Get("year")
//
//	if title == "" || description == "" || year == "" {
//		http.Error(w, "Все поля обязательны", http.StatusBadRequest)
//		return
//	}
//
//	_, err = db.Exec("INSERT INTO movies (title, description, year) VALUES ($1, $2, $3)",
//		title, description, year)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
//}

func main() {

	fmt.Println("movie-service started")
	http.HandleFunc("/api/movies", movieHandler)
	http.HandleFunc("/api/movies/tags", moviesTagsHandler)
	http.HandleFunc("/api/movies/genres", moviesGenresHandler)
	http.ListenAndServe(":8080", nil)
}
