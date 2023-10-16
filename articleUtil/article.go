package articleUtil

import (
	"app/dataAccess"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type articles struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

type IDResponse struct {
	ID int64 `json:"id"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func handleError(w http.ResponseWriter, err error, status int) {
	response := Response{err.Error(), status}
	fmt.Println(err.Error())
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

func handleResponse(w http.ResponseWriter, response interface{}, status int) {
	response = Response{response, status}
	fmt.Println(response)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	posts, err := dataAccess.GetMany("SELECT * FROM articles order by dateCreated desc", reflect.TypeOf(articles{}), nil)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, posts, http.StatusOK)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	article, err := dataAccess.GetByID(id, reflect.TypeOf(articles{}))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, article, http.StatusOK)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article articles
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO articles(title, body, dateCreated, dateUpdated) VALUES(?, ?, NOW(), NOW())"
	args := []interface{}{article.Title, article.Body}

	id, err := dataAccess.PrepareAndExecute(query, args)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, IDResponse{ID: id}, http.StatusCreated)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var article articles
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	err = dataAccess.UpdateById(article, id)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, "Object updated", http.StatusAccepted)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := "DELETE FROM articles WHERE id = ?"
	args := []interface{}{id}

	_, err := dataAccess.PrepareAndExecute(query, args)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, "Deleted", http.StatusAccepted)
}
