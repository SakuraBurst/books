package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
	"github.com/SakuraBurst/books.git/main/repository"
	"github.com/gorilla/mux"
)

type BookControler struct {
	Repository repository.BookRepository
}

func (c BookControler) GetBooks(rw http.ResponseWriter, req *http.Request) {
	books := make([]models.Book, 0)
	err := c.Repository.GetBooksFromDatabase(&books, rw)
	if err != nil {
		SendErrorMessage(rw, err)
	} else {
		js, _ := json.Marshal(books)
		rw.Write(js)
	}
}
func (c BookControler) GetBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	book, er := c.Repository.GetBookFromDatabase(rw, vars["id"])
	if er != nil {
		resp.Encode(er)
	} else {
		resp.Encode(book)
	}

}

func (c BookControler) AddBook(rw http.ResponseWriter, req *http.Request) {
	decode := json.NewDecoder(req.Body)
	book := models.Book{}

	decode.Decode(&book)
	if book.IsBookValid() {
		c.Repository.WriteBookToTheDatabase(rw, book)
		c.GetBooks(rw, req)
	} else {
		SendErrorMessage(rw, nil)
	}
}

func (c BookControler) UpdateBook(rw http.ResponseWriter, req *http.Request) {
	var book models.Book
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	decode := json.NewDecoder(req.Body)
	decode.Decode(&book)
	er := c.Repository.UpdateBookFromDatabase(rw, book, vars["id"])
	if er != nil {
		resp.Encode(er)
	} else {
		book, _ = c.Repository.GetBookFromDatabase(rw, vars["id"])
		resp.Encode(book)
	}

}

func (c BookControler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	er := c.Repository.DeleteBookFromDatabase(rw, vars["id"])
	if er != nil {
		SendErrorMessage(rw, nil)
	} else {
		resp.Encode(models.SuccessMessage)
	}

}
