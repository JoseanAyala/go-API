package main

import (
	"app/articleUtil"
	"app/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file, don't throw an error if it doesn't exist
	godotenv.Load()

	router := createRouter()

	okOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173", "https://joseanayala.vercel.app"})
	allowCredentails := handlers.AllowCredentials()
	okMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	okHeaders := handlers.AllowedHeaders([]string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Cache-Control",
		"Content-Length",
		"Content-Type",
		"Cookie",
		"Host",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
	})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(okOrigins, allowCredentails, okMethods, okHeaders)(router)))
}

func createRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.LogRequests)

	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/articles", articleUtil.GetArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", articleUtil.GetArticleById).Methods("GET")

	r.Handle("/articles/create",
		middleware.EnsureValidToken()(http.HandlerFunc(articleUtil.CreateArticle))).
		Methods("POST")

	r.Handle("/articles/delete/{id}",
		middleware.EnsureValidToken()(http.HandlerFunc(articleUtil.DeleteArticle))).
		Methods("DELETE")

	r.Handle("/articles/edit/{id}",
		middleware.EnsureValidToken()(http.HandlerFunc(articleUtil.UpdateArticle))).
		Methods("PUT")

	return r
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><head><title>Welcome to my API</title><style>body{background-color:#1a1a1a;color:#fff}h1,p,li{color:#fff}</style></head><body><h1>Welcome to my API</h1><p>This is a RESTful API built with Go and Gorilla Mux.</p><p>Endpoints:</p><ul><li>GET /articles - Get all articles</li><li>GET /articles/{id} - Get an article by ID</li><li>POST /articles - Create a new article</li><li>PUT /articles/{id} - Update an article by ID</li><li>DELETE /articles/{id} - Delete an article by ID</li></ul></body></html>")
}
