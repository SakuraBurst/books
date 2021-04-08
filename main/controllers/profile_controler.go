package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SakuraBurst/books.git/main/helpers/crypt"
	"github.com/SakuraBurst/books.git/main/models"
)

func (c Controler) Registration(rw http.ResponseWriter, req *http.Request) {
	var user models.User
	json.NewDecoder(req.Body).Decode(&user)
	if !user.IsValid() {
		SendErrorMessage(rw, errors.New("invalid data"), http.StatusUnprocessableEntity)
	} else {
		user.Password = crypt.CryptPass([]byte(user.Password))
	}

}
