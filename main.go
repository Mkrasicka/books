package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
//         "encoding/json"
//         "github.com/gorilla/mux"
//         "log"
//         "net/http"
// )

//         "log"

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{}

func main() {
	log.Println("Starting server...")
	router := mux.NewRouter()

	router.HandleFunc("/books", handleGetBooks)
	router.HandleFunc("/newbook", handleCreateBook)
	http.ListenAndServe(":3000", router)
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {

	var books = []Book{
		{ID: "1", Title: "Educated", Author: "Tara Westover", Quantity: 3},
		{ID: "2", Title: "Myth of Normal", Author: "Tara Westover", Quantity: 3},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
	}

	log.Printf("Returning book data: %v\n", books)

	// fmt.Printf("to jest data type: %T\n", books)
	booksByte, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(booksByte))
}

func handleCreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request for new book")
	var newBook Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new BOOk
	newBook.ID = fmt.Sprintf("%d", len(books)+1)

	books = append(books, newBook)

	// return the new book data to the client
	booksByte, err := json.Marshal(newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// json.NewEncoder(w).Encode(newBook)
	io.WriteString(w, string(booksByte))
}
