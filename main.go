package main

import (
	"fmt"
	"log"
	"net/http"

	blog "blogAPI/blog"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){
	router:= mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/allblogposts", blog.GetAllPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":1337", router))
}

func main(){
	handleRequests()
}
