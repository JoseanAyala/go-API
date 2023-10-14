package main

import (
	"fmt"
	"log"
	"net/http"

	blog "blogAPI/blog"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", blog.GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", blog.GetArticle).Methods("GET")
	router.HandleFunc("/articles", blog.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", blog.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", blog.UpdateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":1337", router))
}

func main() {
	fmt.Println("Rest API")
	handleRequests()
}
