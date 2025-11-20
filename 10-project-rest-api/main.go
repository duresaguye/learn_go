package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Article struct
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// In-memory store
var (
	articles []Article
	nextID   = 1
	mu       sync.Mutex // To protect concurrent access
)

// Get all articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Create a new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	article.ID = nextID
	nextID++
	articles = append(articles, article)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(article)
}

func main() {
	// Pre-populate
	articles = append(articles, Article{ID: 1, Title: "Hello Go", Content: "Building APIs is fun!"})
	nextID++

	// Routes
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getArticles(w, r)
		case "POST":
			createArticle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("REST API running on :8080...")
	http.ListenAndServe(":8080", nil)
}
