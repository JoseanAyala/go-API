package blog

import (
	db "blogAPI/db"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type blogPost struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	DateCreated string `json:"dateCreated"`
	DateModified string `json:"dateModified"`
}

// Gets all blog posts
func GetAllPosts(w http.ResponseWriter, r *http.Request){
	posts, err := db.Query("SELECT * FROM blogPost", reflect.TypeOf(blogPost{}))
	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	
	fmt.Println("All Blog Posts Endpoint Hit")
	json.NewEncoder(w).Encode(posts)
}