package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SakuraBurst/books.git/main/helpers"
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
	var loginJson map[string]interface{}
	decod := json.NewDecoder(req.Body)
	decod.Decode(&loginJson)
	fmt.Println(loginJson)
	user, err := c.Repository.GetOneFromDatabase(loginJson["email"], "email", models.User{})
	encoder := c.responseEncoder(rw)
	if err != nil {
		c.SendErrorMessage(rw, err, http.StatusUnprocessableEntity)
	} else {
		npMap, err := helpers.DeletePasswordField(user)
		if err != nil {
			userResp := models.UserResponse{User: user}
			encoder.Encode(userResp)
		} else {
			userResp := models.UserMapResponse{User: npMap}
			encoder.Encode(userResp)
		}

	}

}
