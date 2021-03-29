package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Title  string
	ID     int
	Author string
	Year   int
}

type Message struct {
	Text   string
	Status string
}

type ErrorInt interface {
	Error() string
}

func (err Message) Error() string {
	return err.Text
}

var books []Book = make([]Book, 0)

func main() {
	router := mux.NewRouter()
	books = append(books,
		Book{Title: "Гарри Повар и филосовское яйцо", ID: 1, Author: "Kora0108", Year: 2012},
		Book{Title: "Гарри Повар и тайная комната", ID: 2, Author: "Kora0108", Year: 2012},
		Book{Title: "Гарри Повар и узник гамазкабана", ID: 3, Author: "Kora0108", Year: 2012},
		Book{Title: "Гарри Повар и кубок борща", ID: 4, Author: "Kora0108", Year: 2012})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3585", router))
}

var errorMessage Message = Message{Text: "something went wrong", Status: "error"}
var successMessage Message = Message{Text: "zaebis", Status: "success"}

func getBooks(rw http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(books)
	if err != nil {
		err := errorMessage
		errorjson, _ := json.Marshal(err)
		rw.Write(errorjson)
	}
	rw.Write(js)
}
func getBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		resp.Encode(errorMessage)
	}
	fmt.Println(vars["id"])
	bookI, er := searchForBookIndex(id)
	if er != nil {
		resp.Encode(er)
	} else {
		resp.Encode(books[bookI])
	}

}

func deleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		resp.Encode(errorMessage)
	}
	fmt.Println(vars["id"])
	bookI, er := searchForBookIndex(id)
	if er != nil {
		resp.Encode(er)
	} else {
		deleteBookByIndex(bookI)
		resp.Encode(successMessage)
	}

}

func deleteBookByIndex(index int) {
	books = append(books[:index], books[index+1:]...)
}

func searchForBookIndex(id int) (int, error) {
	for index, book := range books {
		if book.ID == id {
			return index, nil
		}
	}
	return -1, errorMessage
}
