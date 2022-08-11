package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const Port = ":8000"

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello World!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	json.NewEncoder(w).Encode(Articles)
}

func requestHandler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(Port, nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	requestHandler()

}
