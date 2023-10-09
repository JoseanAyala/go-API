package blogPost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BlogPost struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	DateCreated string `json:"date_created"`
	DateModified string `json:"date_modified"`
}

func getAllPosts(w http.ResponseWriter, r *http.Request){
	blogPosts := []BlogPost{
		{ID: 1, Title: "My First Blog Post", Body: "This is my first blog post", DateCreated: time.Now().UTC().String(), DateModified:  time.Now().UTC().String()},
		{ID: 2, Title: "My Second Blog Post", Body: "This is my second blog post", DateCreated: time.Now().UTC().String(), DateModified:  time.Now().UTC().String()},
	}
	fmt.Fprintf(w, "All Blog Posts Endpoint Hit")
	json.NewEncoder(w).Encode(blogPosts)
}