package main

import (
	"app/articleUtil"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my API</h1>")
	fmt.Fprintf(w, "<p>This is a RESTful API built with Go and Gorilla Mux.</p>")
	fmt.Fprintf(w, "<p>Endpoints:</p>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, "<li>GET /articles - Get all articles</li>")
	fmt.Fprintf(w, "<li>GET /articles/{id} - Get an article by ID</li>")
	fmt.Fprintf(w, "<li>POST /articles - Create a new article</li>")
	fmt.Fprintf(w, "<li>PUT /articles/{id} - Update an article by ID</li>")
	fmt.Fprintf(w, "<li>DELETE /articles/{id} - Delete an article by ID</li>")
	fmt.Fprintf(w, "</ul>")
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/articles", articleUtil.GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", articleUtil.GetArticleById).Methods("GET")
	router.HandleFunc("/articles", articleUtil.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", articleUtil.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", articleUtil.UpdateArticle).Methods("PUT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	HandleRequests()
}
