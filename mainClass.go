package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"json:content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("EndPoint hit: All Artucles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Homepage Endpoint hit")
}
func handleRequest() {

	myRouter := mux.NewRouter().StringSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articless", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequest()
}
