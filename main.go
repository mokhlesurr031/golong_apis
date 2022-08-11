package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const Port = ":8000"

type Article struct {
	Id      string `json:"id"`
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

func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func requestHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticles)

	log.Fatal(http.ListenAndServe(Port, myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	//Articles = []Article{
	//	Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	//	Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	//}
	requestHandler()

}
