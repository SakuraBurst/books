package controllers

import (
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
)

func (c Controler) Registration(rw http.ResponseWriter, req *http.Request) {
	err := c.Repository.WriteToTheDatabase(models.User{}, req.Body)
	encoder := c.responseEncoder(rw)
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		encoder.Encode(err)
	}

}
