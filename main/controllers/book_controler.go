package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
	"github.com/gorilla/mux"
)

func (c Controler) GetBooks(rw http.ResponseWriter, req *http.Request) {
	books := make([]models.Book, 0)
	err := c.Repository.GetBooksFromDatabase(&books)
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		SendErrorMessage(rw, err, http.StatusNotFound)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		js, _ := json.Marshal(books)
		rw.Write(js)
	}
}
func (c Controler) GetBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	book, er := c.Repository.GetBookFromDatabase(vars["id"])
	if er != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		resp.Encode(er)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		resp.Encode(book)
	}

}

func (c Controler) AddBook(rw http.ResponseWriter, req *http.Request) {
	decode := json.NewDecoder(req.Body)
	book := models.Book{}

	decode.Decode(&book)
	if book.IsBookValid() {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		c.Repository.WriteBookToTheDatabase(book)
		c.GetBooks(rw, req)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		SendErrorMessage(rw, nil, http.StatusNotFound)
	}
}

func (c Controler) UpdateBook(rw http.ResponseWriter, req *http.Request) {
	var book models.Book
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	decode := json.NewDecoder(req.Body)
	decode.Decode(&book)
	if book.IsBookValid() {
		er := c.Repository.UpdateBookFromDatabase(book, vars["id"])
		if er != nil {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusNotFound)
			resp.Encode(er)
		} else {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			book, _ = c.Repository.GetBookFromDatabase(vars["id"])
			resp.Encode(book)
		}

	} else {
		rw.Header().Set("Content-Type", "application/json")
		SendErrorMessage(rw, nil, http.StatusNotFound)
	}

}

func (c Controler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	er := c.Repository.DeleteBookFromDatabase(vars["id"])
	if er != nil {
		rw.Header().Set("Content-Type", "application/json")
		SendErrorMessage(rw, nil, http.StatusNotFound)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		resp.Encode(models.SuccessMessage)
	}

}
