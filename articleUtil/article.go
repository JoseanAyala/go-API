package articleUtil

import (
	"app/dataAccess"
	"app/types"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func handleError(w http.ResponseWriter, err error, status int) {
	response := types.Response{Data: err.Error(), Status: status}
	fmt.Println("Response: ", response)
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

func handleResponse(w http.ResponseWriter, response interface{}, status int) {
	response = types.Response{Data: response, Status: status}

	fmt.Println("Response: ", response)
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
	posts, err := dataAccess.GetMany("SELECT * FROM articles order by dateCreated desc", reflect.TypeOf(types.Articles{}), nil)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, posts, http.StatusOK)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	article, err := dataAccess.GetByID(id, reflect.TypeOf(types.Articles{}))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, article, http.StatusOK)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article types.Articles
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

	handleResponse(w, types.IDResponse{ID: id}, http.StatusCreated)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var article types.Articles
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
