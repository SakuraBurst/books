package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
	"github.com/gorilla/mux"
)

func (c Controler) GetBooks(rw http.ResponseWriter, req *http.Request) {
	books := make([]models.InstanseMaker, 0)
	err := c.Repository.GetAllFromDatabase(&books, models.Book{})
	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusOK)
		js, _ := json.Marshal(books)
		rw.Write(js)
	}
}
func (c Controler) GetBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	book, er := c.Repository.GetOneFromDatabase(vars["id"], "", models.Book{})
	if er != nil {
		rw.WriteHeader(http.StatusNotFound)
		resp.Encode(er)
	} else {
		rw.WriteHeader(http.StatusOK)
		resp.Encode(book)
	}

}

func (c Controler) AddBook(rw http.ResponseWriter, req *http.Request) {
	err := c.Repository.WriteToTheDatabase(&models.Book{}, req.Body)

	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusNotFound)

	} else {
		rw.WriteHeader(http.StatusOK)
		c.GetBooks(rw, req)
	}
}

func (c Controler) UpdateBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	err := c.Repository.UpdateFromDatabase(models.Book{}, req.Body, vars["id"])
	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusOK)
		book, _ := c.Repository.GetOneFromDatabase(vars["id"], "", models.Book{})
		resp.Encode(book)
	}
}

func (c Controler) DeleteBook(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp := json.NewEncoder(rw)
	er := c.Repository.DeleteFromDatabase(vars["id"], "books")
	if er != nil {
		c.SendErrorMessage(rw, nil, http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusOK)
		resp.Encode(models.SuccessMessage)
	}

}
