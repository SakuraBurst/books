package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
)

func (c Controler) Registration(rw http.ResponseWriter, req *http.Request) {
	err := c.Repository.WriteToTheDatabase(models.User{}, req.Body)
	encoder := c.responseEncoder(rw)
	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusUnprocessableEntity)
	} else {
		encoder.Encode(models.SuccessMessage)
	}

}

func (c Controler) Login(rw http.ResponseWriter, req *http.Request) {
	var loginJson map[string]string
	decod := json.NewDecoder(req.Body)
	decod.Decode(&loginJson)
	fmt.Println(loginJson)
	userInt, err := c.Repository.GetOneFromDatabase(loginJson["email"], "email", models.User{})
	user := userInt.(models.User)
	encoder := c.responseEncoder(rw)
	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusUnprocessableEntity)
	} else {
		if user.ComparePassword([]byte(loginJson["password"])) {
			user.DeletePasswordField()
			userResp := models.UserResponse{User: user}
			encoder.Encode(userResp)
		} else {
			err = errors.New("логин или пароль введены неверно")
			c.SendErrorMessage(rw, err, http.StatusUnprocessableEntity)
		}

	}

}
