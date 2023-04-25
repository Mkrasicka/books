package main

import (

	"encoding/json"
	"fmt"
	"io"
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

type Books struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", handleGetBooks)
	http.ListenAndServe(":3000", router)
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books = []Books{
		{ID: "1", Title: "Educated", Author: "Tara Westover", Quantity: 3},
		{ID: "2", Title: "Myth of Normal", Author: "Tara Westover", Quantity: 3},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
	}
	booksByte, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println(string(booksByte))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(booksByte))

}
