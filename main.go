package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Author struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {

	books = append(books, Book{ID: "123", Title: "Title 1", Author: &Author{Title: "Mr", Name: "John Doe"}})
	books = append(books, Book{ID: "456", Title: "Title 2", Author: &Author{Title: "Mrs", Name: "Jane Smith"}})

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/test", printHelloWorld)
	r.HandleFunc("/api/books", getBooks)
	r.HandleFunc("/api/books/{id}", getBook)
	log.Fatal(http.ListenAndServe(":8080", r))
}
