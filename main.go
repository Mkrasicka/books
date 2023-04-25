package main

// import "github.com/gorilla/mux"
// import (
//         "encoding/json"
//         "github.com/gorilla/mux"
//         "log"
//         "net/http"
// )

type Books struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func main() {
	var books = []books{
		{ID: "1", Title: "Educated", Author: "Tara Westover", Quantity: 3},
		{ID: "2", Title: "Myth of Normal", Author: "Tara Westover", Quantity: 3},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
	}

}
