package blog

import (
	db "blogAPI/db"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

// Gets all blog posts
func GetArticles(w http.ResponseWriter, r *http.Request) {
	posts, err := db.Query("SELECT * FROM blogPost", reflect.TypeOf(Article{}))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(posts)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	query := "SELECT * FROM blogPost WHERE id = " + id
	articles, err := db.Query(query, reflect.TypeOf(Article{}))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO blogPost(title, body, dateCreated, dateUpdated) VALUES('?', '?', NOW(), NOW())"
	args := []interface{}{article.Title, article.Body}

	id, err := db.PrepareAndExecute(query, args)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode([]interface{}{id})
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	var args []interface{}
	query := "UPDATE blogPost SET "

	if article.Title != "" {
		query += "title = ?, "
		args = append(args, article.Title)
	}

	if article.Body != "" {
		query += "body = ?, "
		args = append(args, article.Body)
	}

	query += "dateUpdated = NOW() WHERE id = ?;"
	args = append(args, article.Id)

	_, err = db.PrepareAndExecute(query, args)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := "DELETE FROM blogPost WHERE id = ?"
	args := []interface{}{id}

	_, err := db.PrepareAndExecute(query, args)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func handleError(w http.ResponseWriter, err error, status int) {
	fmt.Println(err.Error())
	json.NewEncoder(w).Encode(err.Error())
	w.WriteHeader(status)
}
